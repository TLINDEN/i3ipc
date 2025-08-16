package main

/*
 Retrieve a list of current available workspaces
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

	workspaces, err := ipc.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	repr.Println(workspaces)
}
