package main

import "fmt"

func main() {
	num := make(chan int)
	signal := make(chan struct{})
	go receive(signal, num)
	send(num)

	signal <- struct{}{}

}

func send(num chan<- int) {
	num <- 1
	num <- 2
	num <- 3
	// close(num)
}

func receive(signal <-chan struct{}, num <-chan int) {
	// for v := range num {
	// 	// v, ok := <-num
	// 	// if ok {
	// 	fmt.Printf("%d\n", v)
	// 	// }
	// }
	for {
		select {
		case v := <-num:
			fmt.Println(v)
		case <-signal:
			return
		default:
			fmt.Println("as")
		}
	}
}
