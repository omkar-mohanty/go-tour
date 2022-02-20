package gorutines

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s : %d\n", from, i)
	}
}

func Goro() {
	//direct call to f
	f("direct")

	//gorutine call to f
	go f("gorutine")

	//gorutine closure call
	go func(msg string) {
		fmt.Printf("%s\n", msg)
	}("going")

	//Waiting for gorutines to finish
	time.Sleep(time.Second)

	//Indicating that gorutines have finished executing
	fmt.Println("done")
}
