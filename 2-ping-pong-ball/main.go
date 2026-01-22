package main

import (
	"sync"

	"github.com/frkntplglu/go-concurrency/pingpong"
)

func main() {
	pp := pingpong.New()

	go func() {
		pp.Ball <- 0
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go pp.Play("Player 1", &wg)

	go pp.Play("Player 2", &wg)

	wg.Wait()
}
