# go-tzot

## Overview

This repository provides the following program to utilize timezone offset transitions:

- Go package `"github.com/Jumpaku/go-tzot"`.
- CLI tool `tzot` to generate a slim Go API.

This module is based on the IANA TZ database through a wrapper repository https://github.com/Jumpaku/tz-offset-transitions .

## Go package "github.com/Jumpaku/go-tzot"

### Installation

```shell
go get "github.com/Jumpaku/go-tzot@latest"
```

### Example

```go
package main

import (
	"fmt"
	"github.com/Jumpaku/go-tzot"
)

func main() {
	fmt.Println(tzot.GetTZVersion())
	for _, zoneID := range tzot.AvailableZoneIDs() {
		tzot.GetZone(zoneID) // Got tzot.Zone object
	}
}
```

### Available API

```go
package tzot

import (
	"time"
)

type Zone struct {
	ID          string
	Transitions []Transition
}

type Transition struct {
	When         time.Time
	OffsetBefore time.Duration
	OffsetAfter  time.Duration
}

func GetTZVersion() string

func AvailableZoneIDs() []string

func GetZone(zoneID string) Zone
```


## CLI tool tzot

### Installation

#### Using `go install`

```shell
go install "github.com/Jumpaku/go-tzot@latest"
```

#### Using `go generate`

```go
package main

//go:generate go run "github.com/Jumpaku/go-tzot/cmd/tzot" gen -package=examples -output-path=tzot.go Asia/Tokyo Pacific/Pago_Pago Europe/Zurich Zulu
```

```shell
go generate ./...
```


### Generated API

The generated API is similar to the above but available zone IDs are narrowed to the specified ones with command line arguments to `tzot`.

### Usage

#### tzot gen

```
Generates Go code to handle timezone offset transitions for specified timezone IDs.

Usage:
    $ <program> gen [<option>|<argument>]... [-- [<argument>]...]


Options:
    -all[=<boolean>], -a[=<boolean>]  (default=false):
        Generates Go code for all timezone IDs if true.

    -help[=<boolean>], -h[=<boolean>]  (default=false):
        Shows description of this subcommand.

    -output-path=<string>, -o=<string>  (default=""):
        Specifies output path of gen subcommand. If not specified, stdout is used.

    -package=<string>, -p=<string>  (default="tzot_gen"):
        Specifies package that output API belongs to.


Arguments:
    [0:] [<timezone_id_list:string>]...
        specifies timezone IDs for which Go code is generated.

```

#### tzot list

```
Lists of all available timezone IDs.

Usage:
    $ <program> list [<option>]...


Options:
    -help[=<boolean>], -h[=<boolean>]  (default=false):
        Shows description of this subcommand.

```


