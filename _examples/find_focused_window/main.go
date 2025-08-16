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
			focused.Name, focused.Id, focused.Geometry.Width, focused.Geometry.Height)
	}
}
