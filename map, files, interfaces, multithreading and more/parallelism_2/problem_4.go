package main

import "fmt"

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator1(arguments, done)
	for i := 0; i < 10; i++ {
		arguments <- 1
	}
	close(done)
	fmt.Println(<-result)
}

func calculator1(arguments <-chan int, done <-chan struct{}) <-chan int {
	cn := make(chan int)
	sum := 0
	go func() {
		defer close(cn)
		for {
			select {
			case n := <-arguments:
				sum += n
			case <-done:

				cn <- sum
				return
			}
		}
	}()
	return cn
}
