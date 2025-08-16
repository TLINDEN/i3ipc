package main

/*
 send a  tick with an arbitrary  payload, supply some argument  to the
 command, which  will then  be attached  to the  tick. Use  the events
 example to  retrieve the tick (command  out the Window event  type to
 better see it)
*/

import (
	"fmt"
	"log"
	"os"

	"github.com/tlinden/swayipc"
)

func main() {
	payload := "send_tick.go"

	if len(os.Args) > 1 {
		payload = os.Args[1]
	}

	ipc := swayipc.NewSwayIPC()

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ipc.Close()

	err = ipc.SendTick(payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent tick with payload '%s'.\n", payload)
}
