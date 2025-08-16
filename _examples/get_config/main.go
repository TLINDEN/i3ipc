package main

/*
 Retrieve the user config
*/

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

	config, err := ipc.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
