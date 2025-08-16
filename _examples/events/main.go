package main

/*
   Demonstrate subscribing to events
*/

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alecthomas/repr"
	"github.com/tlinden/swayipc"
)

// Event callback function, needs to implement each subscribed events,
// fed to it as RawResponse
func ProcessTick(event *swayipc.RawResponse) error {
	var err error
	switch event.PayloadType {
	case swayipc.EV_Tick:
		ev := &swayipc.EventTick{}
		err = json.Unmarshal(event.Payload, &ev)
		repr.Println(ev)
	case swayipc.EV_Window:
		ev := &swayipc.EventWindow{}
		err = json.Unmarshal(event.Payload, &ev)
		repr.Println(ev)
	default:
		return fmt.Errorf("received unsubscribed event %d", event.PayloadType)
	}

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ipc := swayipc.NewSwayIPC()

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ipc.Close()

	_, err = ipc.Subscribe(&swayipc.Event{
		Tick:   true,
		Window: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	ipc.EventLoop(ProcessTick)
}
