package main

import (
	"log"

	"github.com/alecthomas/repr"
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

	repr.Println(tree)
}
