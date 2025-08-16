package main

/*
 Retrieve the wm version
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

	version, err := ipc.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	repr.Println(version)
}
