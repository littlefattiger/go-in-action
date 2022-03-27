package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("real GOMAXPROCS", runtime.GOMAXPROCS(-1))
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutings")

	go func() {
		defer wg.Done()
		for c := 0; c < 3; c++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
				time.Sleep(1)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for c := 0; c < 3; c++ {
			for char := 'A'; char < 'A'+26; char++ {

				fmt.Printf("%c ", char)
				time.Sleep(1)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nTerminate Goroutings")

}
