package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string, 1)

	channel <- "message"
	close(channel)

	count := 0

	for {
		select {
		case msg := <-channel:
			count += 1
			fmt.Println(fmt.Sprintf("Message: %s, Count: %d", msg, count))

		case <-time.After(time.Second):
			fmt.Println("Timeout!")

		default:
			fmt.Println("Default")
			channel <- "can I?"
		}
		time.Sleep(time.Second)
	}
}

