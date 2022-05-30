# prom

A simple project management tool written in Go.

## Commands

### init

Initialises a new project. If no name is provided a project is created in the current directory with the name of the directory.

```sh
prom init [name]
```

### close

Closes an open project. If no name is provided the project registered in the current directory is closed.

```sh
prom close [name]
```

## Config

Set the stale project directory:

```sh
prom config set stale_dir <path>
```
