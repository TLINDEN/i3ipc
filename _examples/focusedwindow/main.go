package main

import (
	"fmt"
	"log"

	"github.com/tlinden/i3ipc"
)

func main() {
	ipc := i3ipc.NewI3ipc("SWAYSOCK")

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}

	tree, err := ipc.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	focused := tree.FindFocused()

	if focused != nil {
		fmt.Printf("focused node: %s\n  id: %d\n  Geometry: %dx%d\n",
			focused.Name, focused.Id, focused.Geometry.Width, focused.Geometry.Height)
	}
}
