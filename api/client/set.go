package client

import (
	"fmt"

	Cli "github.com/docker/docker/cli"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/docker/pkg/units"
	"github.com/docker/docker/runconfig"
)

func (cli *DockerCli) CmdSet(args ...string) error {
	cmd := Cli.Subcmd("set", []string{"CONTAINER [CONTAINER...]"}, "Update properties of one or more containers", true)
	flBlkioWeight := cmd.Uint16([]string{"-blkio-weight"}, 0, "Block IO (relative weight), between 10 and 1000")
	flCPUPeriod := cmd.Int64([]string{"-cpu-period"}, 0, "Limit CPU CFS (Completely Fair Scheduler) period")
	flCPUQuota := cmd.Int64([]string{"-cpu-quota"}, 0, "Limit CPU CFS (Completely Fair Scheduler) quota")
	flCpusetCpus := cmd.String([]string{"#-cpuset", "-cpuset-cpus"}, "", "CPUs in which to allow execution (0-3, 0,1)")
	flCpusetMems := cmd.String([]string{"-cpuset-mems"}, "", "MEMs in which to allow execution (0-3, 0,1)")
	flCPUShares := cmd.Int64([]string{"c", "-cpu-shares"}, 0, "CPU shares (relative weight)")
	flMemoryString := cmd.String([]string{"m", "-memory"}, "", "Memory limit")
	flMemorySwap := cmd.String([]string{"-memory-swap"}, "", "Total memory (memory + swap), '-1' to disable swap")
	cmd.Require(flag.Min, 1)

	cmd.ParseFlags(args, true)

	var flMemory int64
	if *flMemoryString != "" {
		parsedMemory, err := units.RAMInBytes(*flMemoryString)
		if err != nil {
			return err
		}
		flMemory = parsedMemory
	}

	var MemorySwap int64
	if *flMemorySwap != "" {
		if *flMemorySwap == "-1" {
			MemorySwap = -1
		} else {
			parsedMemorySwap, err := units.RAMInBytes(*flMemorySwap)
			if err != nil {
				return err
			}
			MemorySwap = parsedMemorySwap
		}
	}

	hostConfig := runconfig.HostConfig{
		BlkioWeight: *flBlkioWeight,
		CpusetCpus:  *flCpusetCpus,
		CpusetMems:  *flCpusetMems,
		CPUShares:   *flCPUShares,
		Memory:      flMemory,
		MemorySwap:  MemorySwap,
		CPUPeriod:   *flCPUPeriod,
		CPUQuota:    *flCPUQuota,
	}

	names := cmd.Args()
	var errNames []string
	for _, name := range names {
		_, _, err := readBody(cli.call("POST", "/containers/"+name+"/set", hostConfig, nil))
		if err != nil {
			fmt.Fprintf(cli.err, "Error trying to set properties of container (%s): %s\n", name, err)
			errNames = append(errNames, name)
		} else {
			fmt.Fprintf(cli.out, "%s\n", name)
		}
	}
	if len(errNames) > 0 {
		return fmt.Errorf("Error: failed to set containers: %v", errNames)
	}

	return nil
}
