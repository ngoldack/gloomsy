# gloomsy

![Project Status](https://img.shields.io/badge/status-WIP-yellow.svg)
[![Go Reference](https://pkg.go.dev/badge/ngoldack/gloomsy.svg)](https://pkg.go.dev/ngoldack/gloomsy)
[![Go Report Card](https://goreportcard.com/badge/github.com/ngoldack/gloomsy)](https://goreportcard.com/report/github.com/ngoldack/gloomsy)
[![Maintainability](https://api.codeclimate.com/v1/badges/4c07a687addab270824c/maintainability)](https://codeclimate.com/github/ngoldack/gloomsy/maintainability)
[![ci](https://github.com/ngoldack/gloomsy/actions/workflows/ci.yaml/badge.svg)](https://github.com/ngoldack/gloomsy/actions/workflows/ci.yaml)

gloomsy contains a collection of tools for creating a GLOOMSY-Stack application.

## Install

### By script

Be careful when running scripts from the internet. Always check the content of the script before executing it.

```sh
curl -sSL https://raw.githubusercontent.com/ngoldack/gloomsy/main/scripts/install.sh | sh -
```

### Homebrew

```sh
brew install ngoldack/tap/gloomsy
```

### By source

Installation from source requires a working Go environment. Go version 1.21 or above is required.

```sh
go install github.com/ngoldack/gloomsy@latest
```

### Manually

1. Download the latest binary from the [release page](https://github.com/ngoldack/gloomsy/releases).
2. Copy the binary to a directory in your PATH.

    ```sh
    sudo cp gloomsy /usr/local/bin
    ```

3. Make the binary executable.

    ```sh
    chmod +x gloomsy
    ```

## Usage

### TL;DR

```sh
gloomsy new --name gloomsy-app
```

### Commands

Run `gloomsy help` to get a list of all available commands.
