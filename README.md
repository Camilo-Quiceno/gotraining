# About This Project

Due to work requirements, I needed to learn recursion in Go. During this process, I searched for basic and practical exercises that would help me practice and apply the concepts I was learning. However, after spending some time searching online, I couldn't find anything that suited my needs. As a result, I decided to create my own set of exercises with the help of GitHub Copilot.

In this repository, you will find the exercises I've designed, along with my solutions to each of them.

# Go routines

## 1. basic_go_routines

We want to create two goroutines that print a message to the console after a 1-second delay. At the end, the program should wait for both goroutines to finish before exiting.

### Expected console output:

```bash
Hello from goroutine 1!
Hello from goroutine 2!
```

**Note:** The order in which the messages appear may vary because goroutines run concurrently.

## 2. synchronizing_multiple_goroutines

In this exercise, you'll create three goroutines that print different messages, but you want them to execute in a specific order. You will use a `sync.WaitGroup` to ensure that all goroutines finish before the main function exits.

### **Requirements:**

1. Create three goroutines:
    - **Goroutine 1**: Prints "First!"
    - **Goroutine 2**: Prints "Second!"
    - **Goroutine 3**: Prints "Third!"
2. The goroutines should print their messages in the order of **First, Second, Third** (even though they run concurrently).
3. You should use a `sync.WaitGroup` to ensure that the main function waits for all the goroutines to finish before printing "Main function finished".

### **Expected Output:**

```bash
First!
Second!
Third!
Main function finished
```

The order of the first three lines should always be as shown (since we will synchronize the goroutines), but **the last line** ("Main function finished") will always appear last after all the goroutines have finished executing.

## 3. concurrent_number_printing

Write a program that starts **5 goroutines** to print numbers from 1 to 5. Each goroutine will print a number, but the order in which the numbers are printed should not matter. The main function should wait for all the goroutines to finish before printing `"All goroutines finished!"`.

**Requirements**:

1. Create 5 separate goroutines, each printing one number.
2. Use `sync.WaitGroup` to ensure the main function waits for all goroutines to finish before exiting.
3. **Do not worry about the order of the output**. Each goroutine will print its number independently, so the order of the prints can vary.

**Ouput**

```bash
3
1
5
4
2
All goroutines finished!
```

# Channels

## 1. send_and_receive

### Objective:

Create a program where a goroutine sends a string message to the main goroutine via a channel, and the main goroutine receives and prints that message.

### Requirements:

1. Create a channel that will send and receive a string.
2. Start a goroutine that sends the message `"Hello from goroutine!"` to the main goroutine.
3. The main goroutine should receive the message from the channel and print it.
4. The main goroutine should wait for the message before printing the final message `"Main goroutine received the message!"`.

### Example Output:

```bash
Hello from goroutine!
Main goroutine received the message!
```

## 2. send_multiple_messages

### Objective:

Create a program that uses a **channel** to send multiple messages from one goroutine to the main goroutine.

### Requirements:

1. Create a channel of type `string`.
2. Start a goroutine that sends **5 messages** to the main goroutine, one message at a time.
3. The main goroutine should receive all 5 messages and print each one.
4. After receiving all the messages, the main goroutine should print `"All messages received!"`.
5. Make sure that the main goroutine does not exit before all messages are received.

---

### Example Output:

```bash
mathematica
Copiar código
Message 1
Message 2
Message 3
Message 4
Message 5
All messages received
```

## 3. buffered_channels

### Objective:

Create a program that demonstrates how a **buffered channel** allows for immediate sending of multiple values without waiting for a receiver, even when there’s a delay in processing each value.

### Requirements:

1. Create a **buffered** channel of type `int` with a capacity of 3.
2. In the main function, immediately send the values `1`, `2`, and `3` to the buffered channel, and print a message before and after each send to confirm they don’t block.
3. Start a separate goroutine that will receive and print each value from the channel with a short delay between each receive operation.
4. Ensure that the main function doesn’t exit until all values have been received and printed by the goroutine.

### Expected Console Output:

```bash
Sending 1 to channel...
Sent 1 to channel.
Sending 2 to channel...
Sent 2 to channel.
Sending 3 to channel...
Sent 3 to channel.
--- Receiving from channel ---
Received 1
Received 2
Received 3
```

This output makes it clear that, with a buffered channel, all three values can be sent to the channel in quick succession without blocking. An unbuffered channel would require each sent value to be received before the next send could proceed, showing a delay after each "Sending" message.

## 4. channel_direction

### Objective:

Create a program that demonstrates **channel direction** by defining a function that only sends messages to a channel and another function that only receives messages from a channel. This will help practice channel direction concepts in Go.

### Requirements:

1. Define a function named `sendMessages` that takes a **send-only** channel as an argument and sends 3 strings to the channel.
2. Define another function named `receiveMessages` that takes a **receive-only** channel as an argument and receives and prints all messages sent to the channel.
3. In the `main` function, create an unbuffered channel of type `string` and start both `sendMessages` and `receiveMessages` as goroutines.
4. Ensure that the main function waits until all messages are received before exiting.

### Expected Console Output:

```
Sending: Hello
Sending: World
Sending: GoLang
Received: Hello
Received: World
Received: GoLang
All messages received!
```

This exercise will help you understand how to specify the direction of channels in function parameters, making it clear which functions are responsible for sending or receiving data on the channel.

# Select

## 1. handle_multiple_channels

### Objective:

Create a program that demonstrates the use of the **`select`** statement to handle multiple channels. This will help you practice handling multiple channels concurrently and performing actions based on which channel is ready.

### Requirements:

1. Create two channels of type `int`, one for sending even numbers and another for sending odd numbers.
2. Launch two goroutines:
    - One goroutine should send even numbers (`0, 2, 4, 6, 8`) to the even channel.
    - Another goroutine should send odd numbers (`1, 3, 5, 7, 9`) to the odd channel.
3. Use the `select` statement in the main function to receive and print values from both channels.
4. Ensure that the program terminates after receiving all 10 values, handling both channels concurrently using `select`.

### Expected Console Output:

```
Received from odd channel: 1
Received from odd channel: 3
Received from even channel: 0
Received from even channel: 2
Received from odd channel: 5
Received from odd channel: 7
Received from even channel: 4
Received from odd channel: 9
Received from even channel: 6
Received from even channel: 8
All numbers received!

```

This exercise demonstrates how the `select` statement allows you to wait on multiple channel operations, choosing the first channel that becomes ready to send or receive a value. This is useful for handling multiple channels concurrently in a Go program.

## 2. select_timeout

### Objective:

Create a program that demonstrates the use of a **timeout** with the `select` statement, and includes print statements showing when data is being sent and received from the channel before the timeout occurs. This will help you practice handling multiple channels, timeouts, and observing channel activity.

### Requirements:

1. Create a channel of type `string`.
2. Start a goroutine that simulates a task (e.g., a simple message) that will be sent to the channel after a delay of 3 seconds.
3. In the main function, use the `select` statement to:
    - Wait for the message from the goroutine.
    - If the message is not received within **2 seconds**, print `"Timeout! No message received."`.
4. Add print statements to indicate when data is being sent to and received from the channel before the timeout.
5. Use the `time.After` function to implement the timeout.

### Expected Console Output:

```
Sending message to channel...
Timeout! No message received.

```

# Sync Package

## Mutex

### 1. protecting_shared_data

### Objective:

Create a program that demonstrates the use of a **mutex** to control access to a shared variable from multiple goroutines. This exercise will help you understand how to avoid race conditions when multiple goroutines are updating a shared resource.

### Requirements:

1. Create a shared integer variable, `counter`, initialized to `0`.
2. Start **5 goroutines** where each goroutine will increment the `counter` variable **1000 times**.
3. Use a `sync.Mutex` to protect access to the `counter` variable so that only one goroutine can increment it at a time.
4. After all goroutines finish, print the final value of `counter`.
5. Use `sync.WaitGroup` to ensure the main function waits until all goroutines have completed.

### Expected Console Output:

The final output should be:

```
Final counter value: 5000
```

### 2. protecting_shared_data_mixed_inc

### Objective:

Create a program where multiple goroutines concurrently **increment** and **decrement** a shared counter. This exercise will help you practice using mutexes to handle complex concurrent operations, where different goroutines perform opposite actions on shared data.

### Requirements:

1. Define a shared integer variable, `counter`, initialized to `0`.
2. Define constants in the main function:
    - `numIncreasers`: the number of goroutines that will **increment** the counter.
    - `numDecreasers`: the number of goroutines that will **decrement** the counter.
    - `operationsPerGoroutine`: the number of increments or decrements each goroutine performs.
3. Create two types of goroutines:
    - **Increment Goroutines**: Each increments the counter by **1** in a loop, `operationsPerGoroutine` times.
    - **Decrement Goroutines**: Each decrements the counter by **1** in a loop, `operationsPerGoroutine` times.
4. Use a `sync.Mutex` to safely access the `counter`.
5. Use a `sync.WaitGroup` to ensure the main function waits until all goroutines finish.
6. Print the final value of `counter` after all goroutines complete.

### Expected Console Output:

If you define:

- `numIncreasers = 5`
- `numDecreasers = 3`
- `operationsPerGoroutine = 1000`

The expected final value of `counter` should be:

```
Final counter value: 2000

```

(Explanation: `(5 - 3) * 1000 = 2000`.)

This exercise will help you practice handling concurrent operations with mutexes, ensuring safe modifications of shared resources.

## Atomic

### 1. atomic_increment

### Objective:

Create a program that uses **atomic operations** to increment a shared counter in a safe way. This exercise will help you understand how to use atomic functions from the `sync/atomic` package to modify shared data in a concurrent environment.

### Requirements:

1. Define a shared integer variable, `counter`, initialized to `0`.
2. Create a constant `numIncrements` that defines how many times the counter should be incremented (e.g., 1000).
3. Launch `numIncrements` goroutines that each increment the `counter` by 1 using an **atomic operation**.
4. Use `sync/atomic.AddInt32` or similar functions to atomically increment the counter.
5. Use a `sync.WaitGroup` to ensure the main function waits until all goroutines complete.
6. Print the final value of `counter` after all goroutines complete.

### Expected Console Output:

If you define:

- `numIncrements = 1000`

The expected final value of `counter` should be:

```
Final counter value: 1000
```

This exercise helps you practice the usage of atomic operations to safely modify shared data in a concurrent environment without using locks like mutexes.

### 2. atomic_multiple_inc

### Objective:

Create a program that demonstrates how to safely perform concurrent **increment** and **decrement** operations on a shared counter using **atomic operations**. This exercise will help you practice using atomic functions in more complex scenarios where multiple types of operations are being performed concurrently.

### Requirements:

1. Define a shared integer variable, `counter`, initialized to `0`.
2. Define constants:
    - `numIncrements`: the number of goroutines that will **increment** the counter.
    - `numDecrements`: the number of goroutines that will **decrement** the counter.
    - `operationsPerGoroutine`: the number of increments or decrements each goroutine performs.
3. Create two types of goroutines:
    - **Increment Goroutines**: Each increments the counter by **1** in a loop, `operationsPerGoroutine` times.
    - **Decrement Goroutines**: Each decrements the counter by **1** in a loop, `operationsPerGoroutine` times.
4. Use `sync/atomic` functions (e.g., `atomic.AddInt64`, `atomic.AddInt32`, etc.) to perform atomic operations on the `counter`.
5. Use a `sync.WaitGroup` to ensure the main function waits until all goroutines finish.
6. Print the final value of `counter` after all goroutines complete.
7. **Bonus Challenge**: Ensure the program handles potential race conditions correctly using atomic operations, and that the final counter value is consistent.

### Expected Console Output:

If you define:

- `numIncrements = 5`
- `numDecrements = 3`
- `operationsPerGoroutine = 1000`

The expected final value of `counter` should be:

```
Final counter value: 2000

```

(Explanation: `(5 * 1000) - (3 * 1000) = 2000`.)

This exercise will help you practice atomic operations in a more complex scenario involving multiple types of concurrent operations on shared data.

## SyncOnce

### 1. sync_conditional_variable

### Objective:

Create a program that demonstrates the use of **sync.Once** to execute an initialization function exactly once, even when multiple goroutines attempt to call it concurrently. This exercise will help you practice using `sync.Once` to safely perform one-time actions in concurrent environments.

### Requirements:

1. Define a shared variable, `data`, initialized to an empty string.
2. Create a `sync.Once` variable called `once` to ensure that the initialization function is only called once.
3. Define a function called `initializeData` that sets `data` to `"Initialized!"` and prints `"Data initialized"`.
4. Create a set number of goroutines (e.g., `5`) that each attempt to call `initializeData` using `sync.Once`.
5. Each goroutine should print the value of `data` after attempting to initialize it.
6. Use a `sync.WaitGroup` to ensure that the main function waits until all goroutines complete.

### Expected Console Output:

```
Data initialized
Data: Initialized!
Data: Initialized!
Data: Initialized!
Data: Initialized!
Data: Initialized!

```

### Notes:

- Only one of the goroutines should actually initialize the `data` variable and print `"Data initialized"`.
- This exercise demonstrates how `sync.Once` can be used to ensure a one-time setup in a concurrent environment.

### 2. resource_init_multiple_resources

### Objective:

Create a program that simulates the initialization of a shared resource (e.g., a database connection) that must occur only once, even though it is required by multiple goroutines at different times. This exercise will help you practice using `sync.Once` in a scenario where goroutines are triggered at different intervals and depend on a shared resource.

### Requirements:

1. Define a shared variable, `dbConnection`, initialized to an empty string.
2. Use a `sync.Once` variable called `once` to ensure the database connection is initialized only once.
3. Define a function called `initializeDB` that sets `dbConnection` to `"Database Connected!"` and prints `"Database initialized"`.
4. Create a function, `accessDatabase`, that simulates a goroutine needing access to `dbConnection`:
    - Each time it is called, it should use `once.Do` to attempt initializing `dbConnection`.
    - It should print `"Accessing: Database Connected!"` after initializing the connection or confirming it is initialized.
5. Use a `sync.WaitGroup` to wait until all goroutines finish.
6. Simulate 10 goroutines calling `accessDatabase` at different intervals by adding random delays (e.g., `time.Sleep`) to each goroutine before calling `accessDatabase`.

### Expected Console Output:

The first goroutine to attempt accessing the database will initialize the connection, while the others will access it without reinitializing:

```bash
Database initialized
Accessing: Database Connected!
Accessing: Database Connected!
Accessing: Database Connected!
...
```

### Notes:

- Only one of the goroutines should print `"Database initialized"`.
- This exercise simulates real-world scenarios where a resource, like a database connection, is needed by multiple parts of a program and should only be initialized once, regardless of concurrent access patterns.

### 3. multiple_init_multiple_resources

### **Objective:**

Create a Go program that concurrently initializes multiple shared resources exactly once, ensuring thread-safe initialization even when accessed by numerous goroutines. Utilize `sync.Once` to manage the one-time initialization of each resource, and handle dependencies between resources where necessary.

### **Requirements:**

1. **Shared Resources:**
    - Implement three distinct resources: `database`, `cache`, and `logger`.
    - Each resource should have its own initialization function that sets its state and prints an initialization message.
2. **Synchronization:**
    - Use separate `sync.Once` instances for each resource to guarantee their initialization occurs only once.
    - Ensure that the `cache` initializes only after the `database` has been successfully initialized.
    - The `logger` should initialize independently of the other two resources.
3. **Goroutine Coordination:**
    - Launch 20 goroutines where:
        - 10 goroutines attempt to access the `database`.
        - 5 goroutines attempt to access the `cache`.
        - 5 goroutines attempt to access the `logger`.
    - Each goroutine should attempt to initialize its respective resource and then perform a mock operation (e.g., printing the resource status).
4. **Dependency Handling:**
    - If a goroutine tries to access the `cache` before the `database` is initialized, it must wait until the `database` initialization is complete before proceeding.
5. **Concurrency Safety:**
    - Ensure that all resource accesses and initializations are free from race conditions.
    - Use a `sync.WaitGroup` to wait for all goroutines to finish their operations before the program exits.

### **Expected Console Output:**

```bash
Database initialized

Cache initialized

Logger initialized

Accessing Database: Database Connected!

Accessing Database: Database Connected!

Accessing Cache: Cache Ready!

Accessing Logger: Logger Active!

Accessing Logger: Logger Active!
```

...

*Note:* The initialization messages for each resource (`"Database initialized"`, `"Cache initialized"`, `"Logger initialized"`) should appear only once, regardless of the number of goroutines. The access messages should reflect the initialized state of each resource.

## Sync Pool

### 1. reuse_buffer

### **Objective:**

Understand how to utilize `sync.Pool` in Go to efficiently manage and reuse objects, minimizing memory allocations in concurrent applications.

### **Requirements:**

1. **Define a Reusable Object:**
    - Create a struct named `Buffer` that contains a slice of bytes.
2. **Initialize a `sync.Pool`:**
    - Set up a pool that manages instances of `Buffer`.
3. **Implement a Worker Function:**
    - Write a function that:
        - Retrieves a `Buffer` from the pool.
        - Appends some data to the `Buffer`.
        - Prints the contents of the `Buffer`.
        - Resets the `Buffer` and returns it to the pool.
4. **Concurrent Execution:**
    - Launch multiple goroutines that execute the worker function concurrently

### **Expected Console Output:**

```bash
Worker 1: Buffer data - Hello from worker 1
Worker 2: Buffer data - Hello from worker 2
```

### **Notes:**

- **Object Reuse:** Observe how `sync.Pool` allows multiple goroutines to reuse `Buffer` instances, reducing the number of allocations.
- **Concurrency Safety:** Ensure that the pool is accessed safely across goroutines without race conditions.
- **Performance Observation:** You can enhance the exercise by measuring memory usage or execution time with and without using `sync.Pool` to see the performance benefits.

---

By completing this exercise, you'll gain practical experience with `sync.Pool`, learning how to optimize resource management and improve the efficiency of your Go applications through effective object reuse.

# Pipeline

## 1. integer_processing

### **Objective:**

Create a Go program that demonstrates a basic pipeline using goroutines and channels. The pipeline will generate integers, process them by squaring each number, and then print the results. This exercise will help you understand how to coordinate multiple stages of processing concurrently.

### **Requirements:**

1. **Number Generator:**
    - Implement a goroutine that generates integers from 1 to 10.
    - Send each integer to the next stage through a channel.
2. **Squaring Processor:**
    - Create a goroutine that receives integers from the generator.
    - Square each received integer.
    - Send the squared result to the next stage through another channel.
3. **Printer:**
    - Develop a goroutine that receives squared integers.
    - Print each squared number to the console.
4. **Pipeline Coordination:**
    - Ensure that all goroutines communicate properly using channels.
    - Close channels appropriately to signal completion and prevent deadlocks.

### **Expected Console Output:**

```bash
Squared Number: 1
Squared Number: 4
Squared Number: 9
Squared Number: 16
Squared Number: 25
Squared Number: 36
Squared Number: 49
Squared Number: 64
Squared Number: 81
Squared Number: 100
```

### **Notes:**

- **Goroutines and Channels:** Utilize goroutines to run each stage of the pipeline concurrently and channels to pass data between them.
- **Channel Closing:** Remember to close channels from the sending side once all data has been sent to avoid deadlocks in receiving goroutines.
- **Synchronization:** Use synchronization techniques, such as `sync.WaitGroup`, if necessary, to ensure the main function waits for all goroutines to finish processing before exiting.
- **Extensibility:** This basic pipeline can be extended by adding more processing stages, such as filtering even numbers or calculating factorials.

## 2. prime_number_pipeline

### **Objective:**

Build a Go application that efficiently identifies prime numbers from a generated list using a multi-stage pipeline. This exercise will deepen your understanding of Go's concurrency primitives, channel synchronization, and pipeline design patterns.

### **Requirements:**

1. **Number Generator:**
    - Create a goroutine that generates integers from 2 up to a specified limit (e.g., 100).
    - Send each integer to the next stage through a channel.
2. **Worker Pool for Prime Checking:**
    - Implement a pool of worker goroutines that receive integers from the generator.
    - Each worker checks if the received number is a prime.
    - If a number is prime, send it to the next stage through another channel.
3. **Prime Collector:**
    - Develop a goroutine that collects prime numbers from the workers.
    - Store the collected prime numbers in a slice.
    - After processing all numbers, print the list of prime numbers.
4. **Pipeline Coordination:**
    - Ensure proper synchronization between all pipeline stages using channels.
    - Close channels appropriately to signal completion and prevent deadlocks.

### **Expected Console Output:**

```bash
Prime Numbers up to 100:
2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97
```

## 3. text_processing_pipeline

### **Objective:**

Create a Go program that processes lines of text through a multi-stage pipeline. The pipeline will read lines from a source, filter lines containing a specific keyword, transform the filtered lines to uppercase, and then print the results. This exercise will help you understand how to build and coordinate concurrent pipelines using goroutines and channels in Go.

### **Requirements:**

1. **Line Generator:**
    - Implement a function that generates lines of text from a predefined slice of strings.
    - Send each line to the next stage through a channel.
2. **Filter Stage:**
    - Create a function that receives lines from the generator.
    - Filter and retain only the lines that contain a specific keyword (e.g., `"ERROR"`).
    - Send the filtered lines to the next stage through another channel.
3. **Transformation Stage:**
    - Develop a function that receives the filtered lines.
    - Convert each line to uppercase.
    - Send the transformed lines to the final stage through another channel.
4. **Printer Stage:**
    - Implement a function that receives the transformed lines.
    - Print each line to the console.
5. **Pipeline Coordination:**
    - Ensure all stages communicate properly using channels.
    - Close channels appropriately to signal completion and prevent deadlocks.
6. **Concurrency:**
    - Run each stage of the pipeline concurrently using goroutines.

### **Expected Console Output:**

Assuming the keyword is `"ERROR"`, the output should display only the lines containing `"ERROR"`, transformed to uppercase.

```bash
ERROR: FAILED TO CONNECT TO DATABASE
ERROR: NULL POINTER EXCEPTION
```

### **Sample Data:**

```go
lines := []string{
    "INFO: Starting the application",
    "DEBUG: Initializing modules",
    "ERROR: Failed to connect to database",
    "INFO: Application running",
    "ERROR: Null pointer exception encountered",
    "DEBUG: Shutting down modules",
    "INFO: Application terminated",
}
```

### **Notes:**

- **Goroutines and Channels:** Utilize goroutines to run each stage of the pipeline concurrently and channels to pass data between them efficiently.
- **Channel Closing:** Properly close channels from the sending side once all data has been sent to avoid deadlocks in receiving goroutines.
- **Synchronization:** Use synchronization techniques, such as `sync.WaitGroup`, if necessary, to ensure the main function waits for all goroutines to finish processing before exiting.
- **Extensibility:** This basic pipeline can be extended by adding more processing stages, such as logging, writing to a file, or performing additional transformations.