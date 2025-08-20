This is a more practical example. With sway you can move windows to
a "scratchpad", i.e. like iconify it. There may be an official way
to get back such windows, but I didn't find a good one. There's the
"scratchpad show" command, but it doesn't allow you to select a
window, it just shows the next one (and it keeps it in the floating
state).

So, this example program lists all windows currently garaged on the
scratchpad. When called with a windows id, it gets back the window
to the current workspace and gives it focus - thus descratching it.

To add comfort to the process I  added a small script which you can
use as a ui to it. It uses  rofi which makes a handy ui. To use it,
compile descratch  with "go  build", copy  the descratch  binary to
some location within your $PATH and run the script.
