package main

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

// SendRequest to manage a request ignoring the context
// '_' is used so the compiler does not complain about the non usage of the variable
func SendRequest(ctx context.Context, arg interface{}) error {
	return nil
}

// Square computes the power of 2^p
func Square(p float64) float64 {
	return math.Pow(2, p)
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	//From the official go documentation: “A goroutine is a lightweight thread of execution”.
	//Goroutines are lighter than a thread so managing them is comparatively less resource intensive.

	printHeader("Basic Goroutines")

	// > go routine is a process that runs at different context
	// The Main function is itself it is a 'goroutine'

	// Goroutines are managed by Go. Goroutines are scheduled accordingly, and over the
	// avaiable cpu-cores, to allow parallelism and concurrency. Go manages automatically all
	// of this depending on the blocking operations, number of cpus, etc..

	// Create a lambda function (command)
	loopFunc := func(number int, count int) {
		fmt.Printf("Started goroutine: %d\n", number)
		defer fmt.Println("Press any button to continue.")
		defer fmt.Printf("Finished goroutine: %d\n", number)
		for i := 1; i < count; i++ {
			time.Sleep(1000)
			fmt.Printf("%d : %d\n", number, i)
		}
	}

	var count = 10
	// Start goroutine #1
	go loopFunc(1, count)
	// Start goroutine #2
	go loopFunc(2, count)

	// The output mixed the two goroutines resutls
	// This is becaus both are running concurrently (parallel?)

	var input string
	fmt.Println("Wait until all the go routines ends !")
	fmt.Scanf("%s", &input)

	printHeader("Using WaitGroup to sync Goroutines")

	// Create a wait group
	var wg sync.WaitGroup
	// Specify the task that are going to be in the group
	wg.Add(3)
	// Each go routine must recrease this value by calling to
	// 'defer wg.Done()' or at the end of the go routine

	// Create a lambda function (command)
	// Use the reference address to &sync.WaitGroup
	waitTask := func(wg *sync.WaitGroup, number int, count int) {
		fmt.Printf("Started goroutine: %d\n", number)
		// Stack the defer functions
		defer wg.Done() // Substract one task to the WaitGroup
		defer fmt.Printf("Finished goroutine: %d\n", number)
		for i := 1; i < count; i++ {
			time.Sleep(1000)
			fmt.Printf("%d : %d\n", number, i)
		}
	}
	fmt.Printf("%T\n", waitTask)

	// Create three go routines
	for _, j := range []int{1, 2, 3} {
		fmt.Printf("Created task: %d\n", j)
		// IMPORTANT
		// It is needed to pass a reference pointer to sync.WaitGroup
		// and not a value copy
		go waitTask(&wg, j, count)
	}

	fmt.Println("Wait until the group ends!")
	// Wait until the goroutines ends
	wg.Wait()
	fmt.Println("WaitGroup have ended!")

	printHeader("Go Channels")

	// Channels are uses to exchange messages between goroutines
	// Channels are bloking operations, so it waits until a message
	// is received to continue

	/*   Force deadlock

	// By default a channel, holds just one message. Then, this chanel waits
	// until it the message is consumed to store another one. (deadlock)

	// Using buffered channel allows to have stored more than one message
	// Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.*

	// Creates a channel (not buffered)
	c := make(chan int)
	// Following send -> consumer (vicevera) throws a deallock
	// Send message to the channel
	fmt.Println("Send message through the channel")
	c <- 4 // fatal error: all goroutines are asleep - deadlock!
	// Consume incomming message from the channel
	fmt.Println("Consume incomming message from the channel")
	cvalue := <-c
	fmt.Println(cvalue)

	*/

	fmt.Println("Simple not blocking channel")
	// Creates a buffered channel (fixed length queue)
	// Fixed queues does not block the calls unitl it reaches the total size
	b := make(chan int, 2)
	b <- 4 // It doesn't block the channel anymore, since there is more free-slots.
	fmt.Println(<-b)
	// Close the channel
	close(b)

	fmt.Println("Buffered channel")
	// Another example with buffer is create a buffered channel.
	// The channel is no waiting for messages to be consumed unitl the
	// messages reach the the size (by default a channel has a length = 1)
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 Error, buffer limit since it waits for empty slots <= 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// Close the channel
	close(ch)

	printHeader("Channels and goroutines")

	// This semale send a message between the two goroutines
	// Create a wait group
	var g sync.WaitGroup
	g.Add(2)
	d := make(chan string)

	// Launch the first goroutine
	// Waits until it receives a message through the channel
	go func(g *sync.WaitGroup, c chan string) {
		defer g.Done()
		fmt.Println("Consumer: Waiting for a message")
		message := <-c
		fmt.Println("Consumer: Message received: " + message)
	}(&g, d)

	// Launch the second goroutine
	// Sends a message through the channel
	go func(g *sync.WaitGroup, c chan string) {
		defer g.Done()
		fmt.Println("Producer: Send message through the channel")
		c <- "This is a message"
		// NOTE: This lines is executed once the channel releases the (unique) slot,
		// since it is not a buffered channel (length = 1) , the go routine
		// waits until other routine/s consume the message or throws a deadlock
		// if there is no consumer.
		fmt.Println("Producer: Message sent and processed")
	}(&g, d)

	// Wait until the goroutines ends
	fmt.Println("Wait until the group ends!")
	g.Wait()
	fmt.Println("WaitGroup have ended!")

	// CLose the channel
	close(d)

	printHeader("Worker Pool")

	// As an example of implementing a working pool in Go using channels

	// Create the needed buffered-channels (queues)
	// - jobs: stores the jobs (numbers) to process
	// - results: stores the results processed from jobs
	items := 12
	jobs := make(chan float64, items)
	results := make(chan float64, items)

	// Creates a func to receive jobs and process them using channels
	worker := func(jobs chan float64, results chan float64) {
		for job := range jobs {
			fmt.Println("Processed: ", job)
			results <- Square(job)
		}
	}

	// Start workers (parallel or concurrent)
	fmt.Println("Start Workers")
	go worker(jobs, results)
	go worker(jobs, results)
	//...
	go worker(jobs, results)

	// Create the jobs (Add them into the jobs channed)
	fmt.Println("Create Jobs")
	for i := 0; i < items; i++ {
		fmt.Printf("%d", i)
		jobs <- float64(i)
	}
	fmt.Println()

	// Waint until until the last job is processed
	fmt.Println("Print results")
	// Following it gives an error
	//for result := range results {
	//	fmt.Println("Result: ", result)
	//}

	// Extract the results from the results channel as being processed
	for i := 0; i < items; i++ {
		fmt.Println("Result: ", <-results)
	}
	fmt.Println("End Jobs created")

	// Close jobs and results channels
	close(jobs)
	close(results)

	printHeader("Select and channels")

	// Select channels are used for not blocking operations when a
	// go routines is waiting for more that one channel to recevice
	// messages from

	// Create different channels
	sc1 := make(chan string)
	sc2 := make(chan string)
	sc3 := make(chan string)

	// Create a lambda function (command)
	selfunc := func(c chan string, index int, delay int64) {
		// Since the channel is one-length, it can be closed from here, so
		// other goroutines detect if the channel has been closed and there
		// is no more messages no be sent.
		defer close(c)
		fmt.Println("Processing channel ", index)
		time.Sleep(time.Duration(delay))
		c <- fmt.Sprintf("%d : %d\n", index, delay)
	}

	// Execute the goroutines (using different channels)
	go selfunc(sc1, 1, 500)
	go selfunc(sc2, 2, 1000)
	go selfunc(sc3, 3, 200)

	tasks := 3
	// While unitl all the tasks has been processed
	for tasks > 0 {
		select {
		case m, op := <-sc1:
			if op {
				fmt.Printf("%T,%T\n", m, op)
				fmt.Printf("First channel processed (%t): %s\n", op, m)
				tasks--
			}
		case m, op := <-sc2:
			if op {
				fmt.Printf("%T,%T\n", m, op)
				fmt.Printf("Second channel processed (%t): %s\n", op, m)
				tasks--
			}
		case m, op := <-sc3:
			if op {
				fmt.Printf("%T,%T\n", m, op)
				fmt.Printf("Third channel processed (%t): %s\n", op, m)
				tasks--
			}
		default:
			fmt.Println("Waiting for messages")
		}
	}
}
