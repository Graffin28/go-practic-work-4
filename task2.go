package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "начал работу", j)
		results <- j * j
		fmt.Println("worker", id, "завершил работу", j)
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, &wg, jobs, results)

	}

	//send 10 jobs
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)

	for r := 1; r <= 10; r++ {
		fmt.Println("Результат: ", <-results)
	}
	fmt.Println("Все работы завершены")
}
