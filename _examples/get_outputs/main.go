package main

/*
 Retrieve a list of current available outputs
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

	outputs, err := ipc.GetOutputs()
	if err != nil {
		log.Fatal(err)
	}

	repr.Println(outputs)
}
