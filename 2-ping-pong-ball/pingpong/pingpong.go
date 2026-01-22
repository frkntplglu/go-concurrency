package pingpong

import (
	"fmt"
	"sync"
)

type PingPong struct {
	Ball chan int
}

func New() *PingPong {
	b := make(chan int)
	return &PingPong{
		Ball: b,
	}
}

func (p *PingPong) Play(player string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		b, ok := <-p.Ball
		if !ok {
			fmt.Printf("Game over. %s is leaving \n", player)
			return
		}
		b++
		if b == 10 {
			close(p.Ball)
			return
		}
		fmt.Printf("%s hit! Score : %d \n", player, b)
		p.Ball <- b

	}
}
