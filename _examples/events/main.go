package main

/*
   Demonstrate subscribing to events
*/

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alecthomas/repr"
	"github.com/tlinden/i3ipc"
)

// Event callback function, needs to implement each subscribed events,
// fed to it as RawResponse
func ProcessTick(event *i3ipc.RawResponse) error {
	var err error
	switch event.PayloadType {
	case i3ipc.EV_Tick:
		ev := &i3ipc.EventTick{}
		err = json.Unmarshal(event.Payload, &ev)
		repr.Println(ev)
	case i3ipc.EV_Window:
		ev := &i3ipc.EventWindow{}
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
	ipc := i3ipc.NewI3ipc()

	err := ipc.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer ipc.Close()

	_, err = ipc.Subscribe(&i3ipc.Event{
		Tick:   true,
		Window: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	ipc.EventLoop(ProcessTick)
}
