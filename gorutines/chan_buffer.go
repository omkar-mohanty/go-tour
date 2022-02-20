package gorutines

import "fmt"

func ChanBuffer() {
	//This creates a buffered channel
	//By default channels send something iff there is a coresponding
	//recieve in the other end
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "Bye"
	/*
		Since channel is buffered for only 2 valeus sending "Bye2"
		will cause runtime terror
		messages <- "Bye2"
	*/
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
