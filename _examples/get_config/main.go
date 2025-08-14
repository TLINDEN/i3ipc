package main

/*
 Retrieve the user config
*/

import (
	"fmt"
	"log"

	"github.com/tlinden/i3ipc"
)

func main() {
	ipc := i3ipc.NewI3ipc()

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}

	config, err := ipc.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
