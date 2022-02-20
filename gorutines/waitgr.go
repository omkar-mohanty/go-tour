package gorutines

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker : %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker : %d done\n", id)
}
func WaitGr() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			//defer wg.Done() ensures that regardless of how the go closure
			//is popped the underlying gorutine is signalled to be done
			defer wg.Done()
			worker(i)
		}()
	}
	//wg.Wait function makes the parent function wait until all the other
	//gorutines have finished executing
	wg.Wait()
}
