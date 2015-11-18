Docker
======
Checkout original [README](https://github.com/docker/docker/blob/master/README.md)

Changes
========
This repository is fork of original [docker repository](https://github.com/docker/docker) with additional feature of `docker set` command. The code is taken from the [patch](https://github.com/hqhq/docker/commit/14c06fce8b4dd38898f6dee6d8be9a2dfcc9bbf7) and re-based on docker version [v1.9.0](https://github.com/docker/docker/tree/v1.9.0). The reference for `docker set` command is available [here](https://github.com/mangalaman93/docker/blob/474e8a4751bd8174f07e1ab24920cf77a4501302/man/docker-set.1.md). The remote API reference is provided [here](https://github.com/mangalaman93/docker/blob/474e8a4751bd8174f07e1ab24920cf77a4501302/docs/reference/api/docker_remote_api_v1.21.md#set-a-container). For current status of this patch in standard docker rope, check out the [pull request](https://github.com/docker/docker/pull/15078)

Binary
======
* docker binary for Linux-amd64 is available [here](https://github.com/mangalaman93/docker/raw/merge_add_set/bundles/1.9.0/binary/docker-1.9.0)
* To compile from source, make sure to keep the project in the folder `$GOPATH/github.com/docker/docker` instead of `$GOPATH/github.com/mangalaman93/docker`

Docker Golang Client
====================
docker golang client with additional feature of `set` command is available [here](https://github.com/mangalaman93/dockerclient)
