## Day 4: "Sluggish Microservice Simulator" (Service Timeout)

### Scenario
You're working on a critical payment system. The system sends a request to the **"Balance Service"** to check a user's balance. However, this service sometimes responds very late due to network issues or high load.

We don't want users to wait for minutes; if the service doesn't respond **within 2 seconds**, you should safely return an error saying **"Operation timed out"**.

---

### Your Task

1. Write a function named **MockService**.  
   This function should take a `chan string` as a parameter.

2. Inside the function, use `time.Sleep` to simulate a **random delay** (between 0 and 4 seconds).

3. After the sleep ends, send this message to the channel: `Balance Information: 1500 TL`

4. In the `main` function, start this service inside a **goroutine**.

5. Use the `select` structure to listen for two cases:
   - **Case A:** Data arriving from the service.
   - **Case B:** The 2-second timeout expiring via `time.After(2 * time.Second)`.

6. Whichever case occurs first, print the appropriate message to the screen and terminate the program.

---

### Technical Requirements

- **Select Usage:**  
  You must use a `select` block.

- **Timeout:**  
  The channel returned by the `time.After` function must be used as a `case` within `select`.

- **Proper Simulation:**  
  To ensure the service sometimes takes less than 2 seconds (success) and sometimes more than 2 seconds (timeout), you can use `rand.Intn`.

---

### Key Concepts You Should Research

- **`select` structure:**  
  What happens if none of the cases are ready? (Blocking)

- **`time.After` function:**  
  How does this function return a channel (`<-chan Time`) in the background?

- **Default Case:**  
  What is the purpose of using `default` inside `select`?  
  (Don't use it in this assignment, but definitely read about what it does.)

---

### A Critical Thinking Question

If a **timeout** occurs and we terminate the program,  
what happens to the **MockService goroutine** that's still running in the background?

> (The answer to this question will lead us to **day 7's topic: the `context` structure**.  
> For now, just think about it.)

---

With this assignment, you'll see how to build **defensive programming** logic using Go's concurrency tools.

When you've prepared the code, share it and let's review it together! 🚀