package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 5; i++ {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}()
	wg.Wait()
	fmt.Println("Done")

}
