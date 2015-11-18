<!--[metadata]>
+++
title = "set"
description = "The set command description and usage"
keywords = ["proprities, update, dynamically"]
[menu.main]
parent = "smn_cli"
weight=1
+++
<![end-metadata]-->

## set

    Usage: docker set [OPTIONS] CONTAINER [CONTAINER...]

    Set properties of one or more containers

      --blkio-weight=0           Block IO weight (relative weight)
      -c, --cpu-shares=0         CPU shares (relative weight)
      --cpuset-cpus=""           CPUs in which to allow execution (0-3, 0,1)
      --cpuset-mems=""           Mems in which to allow execution (0-3, 0,1). Only effective on NUMA systems.
      -m, --memory=0             Memory limit
      --memory-swap=0            Total memory (memory + swap), '-1' to disable swap
      --cpu-period=0             Limit the CPU CFS (Completely Fair Scheduler) period
      --cpu-quota=0              Limit the CPU CFS (Completely Fair Scheduler) quota

The `docker set` command will update properties of one or more containers
wether these containers are running or not.
