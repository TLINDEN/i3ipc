package main

/*
   This  example toggles  borders  globally for  all windows,  execute
   multiple time to see the toggling.
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

	responses, err := ipc.RunGlobalCommand("border toggle")
	if err != nil {
		repr.Println(responses)
		log.Fatal(err)
	}
}
