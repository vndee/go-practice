package main

import "fmt"

func sum_(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum_(s[:len(s) / 2], c)
	go sum_(s[len(s) / 2:], c)

	x, y := <- c, <- c // receive from c
	fmt.Println(x, y, x + y)

	// buffered channel
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
