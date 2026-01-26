package main

import (
	"fmt"
	"sync"

	"github.com/frkntplglu/go-concurrency/bankaccount"
)

func main() {
	ba := bankaccount.New()

	var wg sync.WaitGroup
	for _ = range 1000 {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			ba.Deposit(1)
		}(&wg)
	}

	wg.Wait()

	fmt.Printf("Final Balance: %d €\n", ba.GetBalance())
}
