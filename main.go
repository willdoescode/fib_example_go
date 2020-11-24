package main

import "fmt"

func main() {
	jobs := make(chan int, 45)
	results := make(chan int, 45)

	go worker(jobs, results)

	for i := 0; i < 45; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 45; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n < 3 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i < n+1; i++ {
		a, b = b, a+b
	}
	return a
}
