# Zim Demo

This repository shows how Zim can be used to build and deploy components in a monorepo.

## Getting Started

First up, install [Zim](https://github.com/fugue/zim). You can grab Linux and MacOS (Darwin) binaries from the [releases page](https://github.com/fugue/zim/releases). Confirm it works on your system by running:

```shell
$ zim --version
zim version 0.3.0, build db38827
```

Next, check out this repository, cd into it, and run the following to have Zim print all components:

```shell
$ zim list components
===========================
NAME       | KIND
===========================
authorizer | go_lambda
base       | cloudformation
fake       | node_lambda
opa_cli    | go_binary
reverse    | python_lambda
```

Try running a build by running the following command. Note, by default you'll need Go, Python, and NodeJS available. Note the "Dockerized Builds" option discussed below obviates this requirement.

```shell
$ zim run build
```

## Dockerized Builds

Builds run on the host machine by default in this setup. To opt into running
builds in Docker in order to target a different platform, uncomment these
lines in each build template under [.zim/templates](./.zim/templates).

```yaml
docker:
  image: cimg/go:1.15.6
```

Having the Docker image defines causes Zim to transparently run the build within a Docker container instead of on the host directly.

## Artifacts

Note that by default all output artifacts are stored in a directory `artifacts`
at the top level of the repository.

## Commands to Try

List all components in the repository:

```bash
$ zim list components
```

List all inputs for a given component:

```bash
$ zim list inputs -c authorizer
```

List all outputs for a given component:

```bash
$ zim list outputs -c authorizer
```
