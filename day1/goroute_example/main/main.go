package main

import (
	"fmt"
	"go_learn/day1/goroute_example/goroute"
)

func main() {
	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(100, 300, pipe)

	sum := <-pipe
	fmt.Println("sum:", sum)
}
