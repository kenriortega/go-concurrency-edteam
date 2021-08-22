package main

import (
	"fmt"
	"time"
)

type account struct {
	name    string
	balance int
}
type backOperation struct {
	amount int
	done   chan struct{}
}

func transfer(amount int, source, dest *account) {
	if source.balance < amount {
		fmt.Printf("☠️ :%s\n", fmt.Sprintf("%v %v", source, dest))
		return
	}
	time.Sleep(time.Second)
	dest.balance += amount
	source.balance -= amount
	fmt.Printf("✔️: %s\n", fmt.Sprintf("%v %v", source, dest))
}
func main3() {
	signal := make(chan struct{})
	transaction := make(chan *backOperation)

	a := account{name: "ad", balance: 500}
	b := account{name: "BE", balance: 900}

	go func() {
		for {
			request := <-transaction
			transfer(request.amount, &a, &b)
			request.done <- struct{}{}
		}
	}()
	for _, v := range []int{300, 300} {
		go func(amount int) {
			requestTransaction := backOperation{amount: amount, done: make(chan struct{})}
			transaction <- &requestTransaction

			signal <- <-requestTransaction.done
		}(v)
	}
	<-signal
	<-signal
}
