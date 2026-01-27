package main

import (
	"fmt"

	"github.com/frkntplglu/go-concurrency/generator"
)

func main() {
	gn := generator.New()

	done := make(chan struct{})
	defer close(done)

	nums := gn.RandomNumber(done)

	for i := 0; i < 10; i++ {
		num := <-nums
		fmt.Printf("Generated number %d. : %d\n", i+1, num)
	}

}
