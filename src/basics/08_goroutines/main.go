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
	printHeader("goroutines")

	// a go routine is a process that runs at different context
	// The Main function is itself it is a 'goroutine'

	// Goroutines are managed by Go. Goroutines are scheduled accordingly, and over the 
	// avaiable cpu-cores, to allow parallelism and concurrency. Go manages automatically all
	// of this depending on the blocking operations, number of cpus, etc..

	// Create a lambda function (command)
	loopFunc := func (number int, count int) {
		fmt.Printf("Started goroutine: %d\n", number)
		defer fmt.Println("Remember: Press any button to continue.")
		defer fmt.Printf("Finished goroutine: %d\n", number)
		for i:=1;i<count;i++ {
			time.Sleep(1000)
			fmt.Printf("%d : %d\n", number, i)
		}
	}

	var count = 10
	// Start goroutine 1
	go loopFunc(1, count)
	// Start goroutine 2
	go loopFunc(2, count)

	// The output mixed the two goroutines resutls
	// This is becaus both are running concurrently (parallel?)

	var input string
	fmt.Println("Wait until all the go routines ends !")
	fmt.Println("Press any button to continue.")
	fmt.Scanf("%s",&input)

	printHeader("Using sync.WaitGroup to sync goroutines")

	// Create a wait group
	var wg sync.WaitGroup
	// Specify the task that are going to be in the group
	wg.Add(3)
	// Each go routine must recrease this value by calling to 
	// 'defer wg.Done()' or at the end of the go routine

	// Create a lambda function (command) 
	// Use the reference address to &sync.WaitGroup
	waitTask := func (wg *sync.WaitGroup, number int, count int) {
		fmt.Printf("Started goroutine: %d\n", number)
		// Stack the defer functions
		defer wg.Done() // Substract one task to the WaitGroup
		defer fmt.Printf("Finished goroutine: %d\n", number)
		for i:=1;i<count;i++ {
			time.Sleep(1000)
			fmt.Printf("%d : %d\n", number, i)
		}
	}
	fmt.Printf("%T\n",waitTask)

	// Create three go routines
	for _, j := range []int{1,2,3} {
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

	printHeader("Channels")

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

	// Simple example with buffer, it buffer the 
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 Error, buffer limit
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// CLose the channel
	close(ch)

	// Creates a buffered channel (fixed length queue)
	// Fixed queues does not block the calls
	b := make(chan int, 2)
	// Send message to the channel 
	fmt.Println("Send message through the channel")
	b <- 4
	// Consume incomming message from the channel
	fmt.Println("Consume incomming message from the channel")
	bvalue := <-b
	fmt.Println(bvalue)

	// Close the channel
	close(b)

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
		fmt.Println("Waiting for a message")
		message := <-c
		fmt.Println("Message received: " + message)
	}(&g, d)

	// Launch the second goroutine
	// Sends a message through the channel
	go func(g *sync.WaitGroup, c chan string) {
		defer g.Done()
		fmt.Println("Send message through the channel")
		c <- "This is a message"
		fmt.Println("Message sent and processed")
	}(&g, d)

	// Wait until the goroutines ends
	fmt.Println("Wait until the group ends!")
	g.Wait()
	fmt.Println("WaitGroup have ended!")

	// CLose the channel
	close(d)

	printHeader("Workers Buffer")

	// As an example of implementing a queue

	// Create the needed queues (collections)
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

	// Create the jobs (Add them into the jobs channed)
	fmt.Println("Create Jobs")
	for i := 0;i < items; i++ {
		fmt.Printf("%d",i)
		jobs <- float64(i)
	}
	fmt.Println()

	// Start workers (parallel or concurrent)
	fmt.Println("Start Workers")
	go worker(jobs, results)
	go worker(jobs, results)
	//...
	go worker(jobs, results)

	close(jobs)

	// Waint until until the last job is processed
	fmt.Println("Print results")
	// Following it gives an error
	//for result := range results {
	//	fmt.Println("Result: ", result)
	//} 
	for i := 0;i < items; i++ {
		fmt.Println("Result: ", <-results)
	} 

	fmt.Println("End Workers")

	close(results)

	printHeader("Select and channels")


	sc1 := make(chan string)
	sc2 := make(chan string)
	sc3 := make(chan string)
	// Create a lambda function (command)
	selfunc := func (c chan string, index int, delay int) {
		time.Sleep(delay)
		c <-fmt.Sprintf("%d : %d\n", index, delay)
	}

	// Execute the goroutines
	go selfunc(sc1, 1, 500)
	go selfunc(sc2, 2, 1000)
	go selfunc(sc3, 3, 200)

	select {
	case m := <-sc1:
		fmt.Println("First channel executes: ", m)
	case m := <-sc2:
		fmt.Println("First channel executes: ", m)
	case m := <-sc3:
		fmt.Println("First channel executes: ", m)
	default:
		fmt.Println("Witing for messages")
	}

	// CLose the channel
	close(sc1)
	close(sc2)
	close(sc3)

	printHeader("Context")

	//ctx = context.Background()

	//go SendRequest(c)
}
