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
	f("direct")
	go f("gorutine")
	go func(msg string) {
		fmt.Printf("%s\n", msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}
