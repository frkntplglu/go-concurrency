# Day 3: Image Processing Simulator

## Scenario
Imagine an application with hundreds of images submitted by users. You need to create a thumbnail for each of these images. Image processing is a heavy task, so you want to limit the number of workers (e.g., 3) to working at the same time, with other jobs waiting in a queue.

## Your Task
1. Create a `Job` struct containing:
   * `ID`: Unique identifier for the job
   * `Payload`: A string (e.g., image name)
2. Create a `Result` struct containing:
   * `JobID`: ID of the completed job
   * `Message`: A string message indicating job completion
3. Create two buffered channels:
   * `jobs`: Channel to hold jobs (Capacity: 10)
   * `results`: Channel to collect completed job results (Capacity: 10)
4. Write a `worker` function that:
   * Continuously waits for jobs from the `jobs` channel
   * Simulates heavy processing using `time.Sleep` (e.g., 500ms) when a job arrives
   * Sends the result to the `results` channel after completion
5. In the `main` function:
   * Start exactly **3 worker goroutines**
   * Use a loop to send **15 jobs** to the `jobs` channel
   * Close the `jobs` channel after sending all jobs
   * Read and print results from the `results` channel

## Technical Requirements

### Worker Count Control
- Exactly **3 goroutines** must be running as workers

### Buffered Channels
- Use channel capacities to allow the main goroutine (sender) to continue adding jobs to the channel even if workers haven't caught up yet

### Synchronization
- Use `sync.WaitGroup` to ensure all workers complete their tasks before the program exits
- This is crucial for waiting for all 3 workers to finish processing

## Key Concepts to Research

### Buffered Channels
- What's the difference with `make(chan T, capacity)`?
- How does buffering affect sender and receiver behavior?

### Worker Pool Pattern
- How do workers share a single channel?
- In Go, when multiple receivers read from one channel, each piece of data goes to only one receiver - this acts as an automatic load balancer

### Channel Closing
- Why is it important for the sender to close the channel when done?
- What happens to receivers when a channel is closed?

## Critical Tip

If workers use `range jobs` loop, they will automatically exit when the channel is closed. However, to ensure all workers have exited before closing the `results` channel, using a `WaitGroup` is the most professional approach.

## Expected Learning Outcome

Upon completing this assignment, you will understand how to efficiently manage resources in a system under heavy load using limited, concurrent workers.