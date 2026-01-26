# Day 5: **"Atomic Banking" (The Bank Account Contest)**

### 🏦 Scenario
Imagine a bank account receiving deposits simultaneously from all over the world. If two different processes (**goroutines**) read the balance and try to update it at the exact same moment, one of the transactions might get lost (a classic **Race Condition**). 

Your mission is to prevent this chaos using **`sync.Mutex`**.

---

### 🎯 Your Task
1.  Create a struct named **`BankAccount`** with a **`balance int`** field.
2.  Implement two methods for this struct:
    * **`Deposit(amount int)`** (To add funds)
    * **`GetBalance() int`** (To safely view the balance)
3.  In the **`main`** function, initialize an account.
4.  Start **1000 goroutines**. Each goroutine must deposit **exactly 1 unit** into the account.
5.  After all goroutines have finished (using **`sync.WaitGroup`**), print the final balance.

