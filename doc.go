/*
Package swayipc can be used to control sway, swayfx and possibly
i3wmwindow managers via a unix domain socket.

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

Example usage:

In this example we retrieve the current focused window:

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

Also take a look into the **_examples** folder for more examples.
*/
package swayipc
