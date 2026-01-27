## Day 6: "Infinite Random Number Factory" (Random Stream Generator)

### Scenario
You are working on a data analysis simulation. The system needs a structure that generates random numbers in the background and pushes them into a channel for as long as they are needed. This structure should allow the main program (`main`) to simply consume data from the channel without needing to know how the numbers are generated.

---

### Your Task

1. Write a function named `RandomNumberGenerator(min, max int)`.
2. The return type of this function must be a **receive-only channel** (`<-chan int`).
3. Inside the function, start a **goroutine**, and within that goroutine, generate random numbers within the specified range using an **infinite loop** and send them to the channel.
4. In the `main` function, call this generator and obtain a channel.
5. Using a loop, **read the first 10 numbers** from this channel and print them to the console.
6. After the operation is finished, think of a mechanism to ensure that the generator does **not cause a resource leak (goroutine leak)**.  
   > _(See Hint)_

---

### Technical Requirements

- **Encapsulation**  
  The creation of the channel and the starting of the goroutine must be done inside the function.

- **Receive-only Channel**  
  The return type of the function must not be `chan int`, but **`<-chan int`**.  
  This prevents the consumer from accidentally sending data into the channel.

- **Non-blocking Start**  
  When the generator is called, it must **return the result (the channel) immediately**, while the number generation continues in the background.

---

### Key Concepts You Should Research

- **Directional Channels**
  - `chan T` → bidirectional  
  - `<-chan T` → receive-only  
  - `chan<- T` → send-only  

- **Goroutine Leak**  
  When the main program stops consuming data from the generator, what happens to the infinite loop inside the generator function?  
  Does that goroutine remain **blocked and stuck**?
