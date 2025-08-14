package main

/*
 Retrieve a list of current available outputs
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

	outputs, err := ipc.GetOutputs()
	if err != nil {
		log.Fatal(err)
	}

	repr.Println(outputs)
}
