package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	direction := 1
	for {
		robotgo.MoveRelative(direction*100, direction*200)
		fmt.Println("Sleeping for 20 seconds.")
		time.Sleep(20 * time.Second)
		// alternate the direction
		direction = direction * -1
	}
}
