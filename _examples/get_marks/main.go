package main

/*
 Retrieve a list of current set marks
*/

import (
	"log"

	"github.com/alecthomas/repr"
	"github.com/tlinden/i3ipc"
)

func main() {
	ipc := i3ipc.NewI3ipc()

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
