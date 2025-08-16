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

	tree, err := ipc.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	workspace := tree.FindCurrentWorkspace()

	fmt.Printf("current workspace: %s\n", workspace)
}
