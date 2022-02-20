package gorutines

import "fmt"

func ChannelTest() {
	//this is to test channels

	//This makes a channel of type string
	messages := make(chan string)
	/*
		A normal channel would not allow values to be sent in the same
		function see chan_buffer.go
		messages <- "ping"
	*/
	go func() {
		//using the chan <- input syntax "ping" is put into the channel
		messages <- "ping"
	}()

	//using the <- chan syntax the value from messages is put into msg
	//by default channels are blocking and the consumer will wait until
	//the producer has put some values in the channel
	msg := <-messages
	fmt.Printf("Message : %s\n", msg)
}
