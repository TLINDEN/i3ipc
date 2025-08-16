package main

/*
 Retrieve a list of running bars
*/

import (
	"fmt"
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

	bars, err := ipc.GetBars()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current bars:")
	repr.Println(bars)

	if len(bars) > 0 {
		fmt.Println("First bar:")
		bar, err := ipc.GetBar(bars[0])
		if err != nil {
			log.Fatal(err)
		}

		repr.Println(bar)
	}
}
