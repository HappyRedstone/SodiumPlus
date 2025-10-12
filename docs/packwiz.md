# Packwiz Installation and Usage

This document explains how to install and use the `packwiz` tool.

## Installation

To install `packwiz`, use the following Go command:

```bash
go install github.com/packwiz/packwiz@latest
```

## Usage

After installation, you need to add the Go binary directory to your `PATH` environment variable. Then you can use `packwiz` commands.

```bash
export PATH="$HOME/go/bin:$PATH"
packwiz --help
