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

	workspace := tree.FindCurrentWorkspace()

	fmt.Printf("current workspace: %s\n", workspace)
}
