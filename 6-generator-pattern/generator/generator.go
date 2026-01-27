package generator

import (
	"math/rand"
	"time"
)

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) RandomNumber(done <-chan struct{}) <-chan int {
	num := make(chan int)
	go func() {
		for {
			select {
			case <-done:
				close(num)
				return
			case num <- rand.Intn(10000):
				time.Sleep(100 * time.Millisecond)
			}

		}
	}()

	return num
}
