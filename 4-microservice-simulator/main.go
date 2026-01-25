package main

import (
	"fmt"
	"time"

	balanceservice "github.com/frkntplglu/go-concurrency/balance-service"
)

func main() {
	bs := balanceservice.New()

	balanceVal := make(chan string, 1)

	go bs.CheckBalance(balanceVal)

	select {
	case result := <-balanceVal:
		fmt.Println(result)
		return
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout err")

	}

}
