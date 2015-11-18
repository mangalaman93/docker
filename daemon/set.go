package daemon

import (
	"github.com/docker/docker/runconfig"
)

func (daemon *Daemon) ContainerSet(name string, hostConfig *runconfig.HostConfig) ([]string, error) {
	var warnings []string

	container, err := daemon.Get(name)
	if err != nil {
		return warnings, err
	}

	warnings, err = daemon.verifyContainerSettings(hostConfig, nil)
	if err != nil {
		return warnings, err
	}

	if err := container.Set(hostConfig); err != nil {
		return warnings, err
	}

	return warnings, nil
}
