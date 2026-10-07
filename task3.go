package main

import (
	"fmt"
	"time"
)

func main() {
	semafor := make(chan struct{}, 5)

	tick := time.Tick(200 * time.Millisecond)

	for i := 1; i < 16; i++ {
		go func() {
			semafor <- struct{}{}
			defer func() { <-semafor }()
			fmt.Println(i)
		}()
		<-tick
	}
}
