package main

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

	tree, err := ipc.GetTree()
	if err != nil {
		log.Fatal(err)
	}

	repr.Println(tree)
}
