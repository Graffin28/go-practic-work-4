package main

import(
	"fmt"
	"sync"
	"os"
	"crypto/md5"
)

func () {
	files := []string{}
	jobs := make(chan string, len(files))

	var wg sync.WaitGroup

	for _, i := range files {
		jobs <- i
	}
	close(jobs)


	for worker := 1; worker <= 2; worker++{

		wg.Add(1)

		go func() {
			defer wg.Done()

			for fileN := range jobs {
				data, err := os.ReadFile(fileN)

				if err != nil {
					fmt.Printf(worker, fileN, err)
					coninue
				}
				H := md5.Sum(data)
				fmt.Printf(worker, fileN, H)
			}
		}()
	}
	wg.Wait()
	fmt.Println("End")
}
