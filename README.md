# go CLI for Litra Glow

```
A CLI tool for controlling Logitech Litra Glow lights

Usage:
  logislim [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Returns a list of compatible Logitech Litra Glow lights
  off         Turns off all available Litra Glow lights
  on          Turns on all available Litra Glow lights

Flags:
  -h, --help   help for logislim
```

# Installation

## Prebuilt binaries

1. Download the archive from the releases page
2. Extract the archive
3. Move the executable to the desired directory
4. Add this directory to the PATH environment variable
5. Verify that you have execute permission on the file

## Build from source

```
go install github.com/raulcodes/logislim@latest
```
