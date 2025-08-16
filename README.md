[![Go Report Card](https://goreportcard.com/badge/github.com/tlinden/swayipc)](https://goreportcard.com/report/github.com/tlinden/swayipc) 
[![Actions](https://github.com/tlinden/swayipc/actions/workflows/ci.yaml/badge.svg)](https://github.com/tlinden/swayipc/actions)
![GitHub License](https://img.shields.io/github/license/tlinden/swayipc)
[![GoDoc](https://godoc.org/github.com/tlinden/swayipc?status.svg)](https://godoc.org/github.com/tlinden/swayipc)

# swayipc - go bindings to control sway (and possibly i3)

Package swayipc can be used to control [sway](https://swaywm.org/),
[swayfx](https://github.com/WillPower3309/swayfx) and possibly
[i3wm](http://i3wm.org/) window managers via a unix domain socket.

## About

swaywm's interprocess  communication (or  ipc) is the  interface sway,
swayfx and i3wm use to  receive commands from client applications such
as  sway-msg.   It  also  features  a  publish/subscribe  mechanism  for
notifying interested parties of window manager events.

swayipc is a go module for  controlling the window manager. This project
is intended to  be useful for general scripting,  and for applications
that interact  with the  window manager  like status  line generators,
notification daemons, and window pagers. It is primarily designed to
work with sway and swayfx, but may also work with i3wm, although I
haven't tested it on i3wm.

The module uses the i3-IPC proctocol as outlined in sway-ipc(7).

For details on how to use the library, see the
[reference documentation](https://godoc.org/github.com/tlinden/swayipc).

## Example usage

In this example we retrieve the current focused window:

```go
package main

import (
        "fmt"
        "log"

        "github.com/tlinden/swayipc"
)

func main() {
        ipc := swayipc.NewSwayIPC()

        err := ipc.Connect()
        if err != nil {
                log.Fatal(err)
        }
        defer ipc.Close()

        tree, err := ipc.GetTree()
        if err != nil {
                log.Fatal(err)
        }

        focused := tree.FindFocused()

        if focused != nil {
                fmt.Printf("focused node: %s\n  id: %d\n  Geometry: %dx%d\n",
                        focused.Name, focused.Id, focused.Geometry.Width,
                        focused.Geometry.Height)
        }
}
```

Also take a look into the **_examples** folder for more examples.

You may take a look at the [tool swaycycle](https://github.com/tlinden/swaycycle)
which is using this module.

## Installation

Execute this to add the module to your project:
```sh
go get github.com/tlinden/swayipc
```

## Acknowledgements

A  couple   of  ideas   have  been  taken   from  the   [i3ipc  python
module](https://github.com/altdesktop/i3ipc-python/),   although  this
one is not just a port of it and has been written from scratch.

## Getting help

Although I'm happy to hear from swayipc users in private email, that's the
best way for me to forget to do something.

In order to report a bug,  unexpected behavior, feature requests or to
submit    a    patch,    please    open   an    issue    on    github:
https://github.com/TLINDEN/swayipc/issues.

## Copyright and license

This software is licensed under the GNU GENERAL PUBLIC LICENSE version 3.

## Authors

T.v.Dein <tom AT vondein DOT org>

## Project homepage

https://github.com/TLINDEN/swayipc

## Copyright and License

Licensed under the GNU GENERAL PUBLIC LICENSE version 3.

## Author

T.v.Dein <tom AT vondein DOT org>

