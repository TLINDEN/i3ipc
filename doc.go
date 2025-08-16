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
*/
package swayipc
