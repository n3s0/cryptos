# cryptos

## Overview

A little command line app I wrote in go to help with CTF challenges that need
hex converted to whatever output needed.

## Usage

Here is the usage of the cli app. This has two flags. One that will take input
and the other that indicates it is to be converted to a string.

```sh
CLI application that will convert hex strings to ascii strings

Usage:
  hex-convert [flags]

Flags:
  -h, --help           help for hex-convert
  -i, --input string   Hex input to convert. Needs to be Hex..
  -s, --to-string      Convert hex to string
```

## Go Install

This is the recommended method. It will build and move the installation to the
`$HOME/go/bin` directory. Where if this directory is in the accounts `PATH`
variable. The `hex-convert` binary can be run from that accounts home directory
instead of making it available to almost everyone.

```sh
go install .
```

## Build

This application can be built within the project directory. It will output a
file named `hex-convert`.

```sh
go build .
```

If it builds successfully it can be moved to something like `/usr/bin/` or
wherever else you decide to move it.

```sh
sudo mv ./hex-convert /usr/bin/
```

## Running

This can be build and then run or it can be run using Go. I normally run it
using go in the project directory.

```sh
go run main.go -i <hex string> -s
```
