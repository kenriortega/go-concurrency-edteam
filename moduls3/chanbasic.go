package main

import "fmt"

func main1() {
	num := make(chan int, 2)
	signal := make(chan struct{})
	go receive(signal, num)
	send(num)

	<-signal
}

func send(num chan<- int) {
	num <- 1
	num <- 2
	num <- 3
}

func receive(signal chan<- struct{}, num <-chan int) {
	fmt.Println(<-num)
	fmt.Println(<-num)
	fmt.Println(<-num)
	signal <- struct{}{}
}
