# Day 2: "Ping-Pong Match" (Ping-Pong Ball)

## Scenario
You will simulate a table tennis (ping-pong) match between two different players (goroutines). A ball (actually an integer or struct) will be sent through a channel, the other player will receive it, increment the hit count, and send the ball back through the same channel.

## Your Task
1. Create a `ball` structure or a simple `int` counter.
2. Start two goroutines (e.g., `player1` and `player2`).
3. These two goroutines must communicate through a single unbuffered channel (`chan int`).
4. When the game starts:
   * Player 1 hits the ball to the other side.
   * Player 2 catches the ball, increments the hit count by 1, prints "Player 2 hit! Score: X" to the screen, and sends the ball back.
   * Player 1 catches the ball, increments the hit count, prints to the screen, and sends it back.
5. When the hit count reaches 10, the game should end and the channel should be closed.

## Technical Requirements
* **Unbuffered Channel**: Must be created without specifying capacity like `make(chan int)`.
* **Synchronization**: Using the channel's natural blocking property, ensure that one player cannot hit before the other hits.
* **Termination**: When the channel is closed (`close`), ensure that goroutines terminate gracefully without causing "panic" (by checking `v, ok := <-ch`).

## Key Concepts You Should Research
* `chan` data type and its creation with `make(chan T)`.
* Sending data through a channel (`ch <- data`) and receiving data (`data := <-ch`).
* Closing a channel (`close(ch)`) and checking for closed channel reads (`ok` idiom).
* Consuming data from a channel using `for range` loop (An alternative method).

## A Critical Question (For You to Think About)
If both players are waiting to receive data from the channel at the same time (Receive) and no one has put the ball into the channel for the first time, what happens? How should you serve the first "serve" in the `main` function to start the game?