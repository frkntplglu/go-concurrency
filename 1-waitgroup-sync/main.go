package main

import (
	"context"
	"time"

	"github.com/frkntplglu/go-concurrency/client"
	"github.com/frkntplglu/go-concurrency/healthchecker"
)

func main() {
	myClient := client.New(10 * time.Second)
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.golang.org",
		"https://www.medium.com",
		"https://www.twitter.com",
		"https://www.apple.com",
		"https://www.amazon.com",
	}
	checker := healthchecker.New(urls, myClient)

	checker.CheckConcurrently(context.Background())

}
