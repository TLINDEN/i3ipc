package main

/*
   This  example  toggles  the   current  terminal  window's  floating
   state.  Execute  repeatedly to  toggle  between  full and  floating
   state. Run this from a terminal window to see the effect.
*/

import (
	"log"

	"github.com/alecthomas/repr"
	"github.com/tlinden/i3ipc"
)

func main() {
	ipc := i3ipc.NewI3ipc()

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// retrieve the sway tree
	tree, err := ipc.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	// find the data for the current window
	focused := tree.FindFocused()

	if focused == nil {
		log.Fatal("no focused window found")
	}

	// finally execute  the given commands on it, you  can use any run
	// command, see sway(5)
	responses, err := ipc.RunCommand(focused.Id, "floating toggle, border toggle")
	if err != nil {
		repr.Println(responses)
		log.Fatal(err)
	}
}
