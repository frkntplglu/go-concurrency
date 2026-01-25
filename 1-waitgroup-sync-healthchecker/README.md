# Day 1: Website Health Check (Status Checker)

### Scenario

You are a system administrator and you have 10–15 websites that need to be checked regularly.
Your task is to verify whether these websites are up and running by sending an HTTP GET request and checking if they return a 200 OK status code.

If you check them sequentially, you must wait for each request to finish before starting the next one, which makes the total execution time unnecessarily long.

👉 To solve this efficiently, you will use concurrency in Go.

### Task

Using the URL list below, write a Go program that sends HTTP GET requests concurrently and prints the result for each website.

```
urls := []string{
    "https://www.google.com",
    "https://www.github.com",
    "https://www.golang.org",
    "https://www.medium.com",
    "https://www.twitter.com",
    "https://www.apple.com",
    "https://www.amazon.com",
}
```

## Technical Requirements

- Each URL check must run in its own goroutine
- The main function must not exit before all goroutines finish
- Use `sync.WaitGroup`
- Each goroutine must print:
  - The URL
  - The returned HTTP status (e.g. `200 OK`)
- Start a timer at the beginning of the program
- Print the total execution time at the end using `time.Since`

## Key Concepts to Research

- Starting goroutines with the `go` keyword
- `sync.WaitGroup`
  - `.Add()`
  - `.Done()`
  - `.Wait()`
- Closure problem in goroutines
  - Why loop variables should be passed as parameters instead of being captured directly
