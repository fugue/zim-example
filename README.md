# Zim Demo

This repository shows how Zim can be used to build and deploy multiple
components within a monorepo.

## Getting Started

Install [Zim](https://github.com/fugue/zim).

Then run `zim run build` to build all components.

## Dockerized Builds

Builds run on the host machine by default in this setup. To opt into running
builds in Docker in order to target a different platform, uncomment these
lines in each build template:

```yaml
docker:
  image: cimg/go:1.15.6
```

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
