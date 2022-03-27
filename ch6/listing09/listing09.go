package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {

	wg.Add(2)

	fmt.Println("Start Goroutings")

	go incCounter(1)
	go incCounter(2)

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nFinal Counter: ", counter)

}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		v := counter
		runtime.Gosched()
		v++

		counter = v
	}
}
