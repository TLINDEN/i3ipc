package main

/*
 Retrieve a list of current set marks
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

	marks, err := ipc.GetMarks()
	if err != nil {
		log.Fatal(err)
	}

	repr.Println(marks)
}
