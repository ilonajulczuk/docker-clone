# docker-clone

Little command line program to create clones of Docker containers.

## Use case:

- you have a container running on your host and it's correctly configured
- you want to run another container just like that, but avoid name / port binding collisions and
having the same (or different) env variables, volumes and resource constraints


## What it does / intends to do:

- get the config of the original container
- create and start copy of it, automatically avoiding collisions, letting you override exactly what you want (defaults from original container).

## Usage

```
$ go run clone.go original_container
```

## TODO:

Most of the project. 

- Collision avoidance
- Overrides
- Interactive questions



