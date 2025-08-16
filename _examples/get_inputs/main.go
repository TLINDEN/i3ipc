package main

/*
 Retrieve a list of current available inputs
*/

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
	defer ipc.Close()

	inputs, err := ipc.GetInputs()
	if err != nil {
		log.Fatal(err)
	}

	repr.Println(inputs)
}
