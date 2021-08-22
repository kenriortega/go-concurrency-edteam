package main

import (
	"fmt"
	"sync"
	"time"
)

type account struct {
	name    string
	balance int
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
func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	wg.Add(2)
	a := account{name: "ad", balance: 500}
	b := account{name: "BE", balance: 900}

	for _, v := range []int{300, 300} {
		go func(amount int) {
			mu.Lock()
			transfer(amount, &a, &b)
			mu.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
}
