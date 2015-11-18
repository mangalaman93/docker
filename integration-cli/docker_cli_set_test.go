package main

import (
	"github.com/go-check/check"
)

func (s *DockerSuite) TestSetContainer(c *check.C) {
	name := "test-set-container"
	dockerCmd(c, "run", "-d", "--name", name, "-m", "300M", "busybox", "true")
	dockerCmd(c, "set", "-m", "500M", name)

	memory, err := inspectField(name, "HostConfig.Memory")
	c.Assert(err, check.IsNil)

	if memory != "524288000" {
		c.Fatalf("Got the wrong memory value, we got %d, expected 524288000(500M).", memory)
	}
}

func (s *DockerSuite) TestSetContainerInvalidValue(c *check.C) {
	name := "test-set-container"
	dockerCmd(c, "run", "-d", "--name", name, "-m", "300M", "busybox", "true")

	_, _, err := dockerCmdWithError("set", "-m", "2M", name)
	if err == nil {
		c.Fatal("[set] should failed if we tried to set invalid value.")
	}
}
