package main

/*
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
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alecthomas/repr"
	"github.com/tlinden/swayipc"
)

func main() {
	// we need a session to sway via IPC
	ipc := swayipc.NewSwayIPC()

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ipc.Close()

	// first, retrieve the whole sway root node
	root, err := ipc.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	// get the hidden scratchpad workspace
	scratch := getScratch(root)

	if len(os.Args) > 1 {
		// called with an arg, consider it to be a container id
		id, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("failed to convert arg %s to integer: %w", os.Args[1], err)
		}

		// switch to it
		retrieveWindow(ipc, scratch, id)
	} else {
		// no args, just list scratched windows
		listScratchedContainer(scratch)
	}
}

func retrieveWindow(ipc *swayipc.SwayIPC, scratch *swayipc.Node, id int) {
	// make sure the id exists
	var exists bool
	for _, con := range scratch.FloatingNodes {
		if con.Id == id {
			exists = true
		}
	}

	if !exists {
		log.Fatalf("no window with id %d exists on the scratchpad", id)
	}

	// scratched windows are floating, so we move it to current
	// workspace, disable the floating state and switch focus to it
	responses, err := ipc.RunContainerCommand(id, "move workspace current", "floating toggle", "focus")
	if err != nil {
		repr.Println(responses)
		log.Fatal(err)
	}

}

func listScratchedContainer(scratch *swayipc.Node) {
	// list the windows
	for _, con := range scratch.FloatingNodes {
		fmt.Printf("%d %s\n", con.Id, con.Name)
	}
}

func getScratch(root *swayipc.Node) *swayipc.Node {
	// the root node only has output nodes, we iterate over them and
	// look for the internal one named__i3. This in turn only has one
	// workspace, the scratchpad workspace, which we return.
	for _, output := range root.Nodes {
		if output.Name == "__i3" {
			return output.Nodes[0]
		}
	}

	return nil
}
