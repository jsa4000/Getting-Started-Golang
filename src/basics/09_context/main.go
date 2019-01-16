package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// sleepRandom function that sleeps a random time unitl it is processed
func sleepRandom(fromFunction string, ch chan int) {
	//defer cleanup
	defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()

	//Perform a slow task
	//Sleep here for random ms
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100
	fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")

	//write on the channel if it was passed in
	if ch != nil {
		ch <- sleeptime
	}
}

// Function that does slow processing with a context
// Note that context is the first argument
func sleepRandomContext(ctx context.Context, ch chan bool) {

	// Cleanup tasks
	// There are no contexts being created here
	// Hence, no canceling needed
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	//Make a channel
	sleeptimeChan := make(chan int)

	//Start slow processing in a goroutine
	//Send a channel for communication
	go sleepRandom("sleepRandomContext", sleeptimeChan)

	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//If context expires, this case is selected
		//Free up resources that may no longer be needed because of aborting the work
		//Signal all the goroutines that should stop work (use channels)
		//Usually, you would send something on channel,
		//wait for goroutines to exit and then return
		//Or, use wait groups instead of channels for synchronization
		fmt.Println("Time to return")
	case sleeptime := <-sleeptimeChan:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}

//A helper function, this can, in the real world do various things.
//In this example, it is just calling one function.
//Here, this could have just lived in main
func doWorkContext(ctx context.Context) {

	//Derive a timeout context from context with cancel
	//Timeout in 150 ms
	//All the contexts derived from this will returns in 150 ms
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

	//Cancel to release resources once the function is complete
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	//Make channel and call context function
	//Can use wait groups as well for this particular case
	//As we do not use the return value sent on channel
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)

	//Use a select statement to exit out if context expires
	select {
	case <-ctx.Done():
		//This case is selected when the passed in context notifies to stop work
		//In this example, it will be notified when main calls cancelFunction
		fmt.Println("doWorkContext: Time to return")
	case <-ch:
		//This case is selected when processing finishes before the context is cancelled
		fmt.Println("sleepRandomContext returned")
	}
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	// http://p.agnihotry.com/post/understanding_the_context_package_in_golang/

	// The context package in go can come in handy while interacting with APIs and
	// slow processes, especially in production-grade systems that serve web requests.
	// Where, you might want to notify all the goroutines to stop work and return. Here
	// is a basic tutorial on how you can use it in your projects with some best practices and gotchas.

	// A way to think about context package in go is that it allows you to pass in a “context” to
	// your program. Context like a 'timeout' or 'deadline' or a channel to indicate stop working and return.
	// For instance, if you are doing a web request or running a system command, it is usually a good idea
	// to have a timeout for production-grade systems. Because, if an API you depend on is running slow, you
	// would not want to back up requests on your system, because, it may end up increasing the load and degrading
	// the performance of all the requests you serve. Resulting in a cascading effect. This is where a timeout
	// or deadline context can come in handy.

	//  Creating context

	// The context package allows creating and deriving context in following ways:

	// context.Background() ctx Context

	// This function returns an empty context. This should be only used at a high level (in main or the
	// top level request handler). This can be used to derive other contexts that we discuss later.

	// ctx, cancel := context.Background()
	// context.TODO() ctx Context

	// This function also creates an empty context. This should also be only used at a high level
	// or when you are not sure what context to use or if the function has not been updated to receive
	// a context. Which means you (or the maintainer) plans to add context to the function in future.

	// ctx, cancel := context.TODO()

	// Interestingly, looking at the code (https://golang.org/src/context/context.go), it is
	// exactly same as background. The difference is, this can be used by static analysis
	// tools to validate if the context is being passed around properly, which is an important
	// detail, as the static analysis tools can help surface potential bugs early on, and can be
	// hooked up in a CI/CD pipeline.

	// Conetxt types
	// context.WithValue(parent Context, key, val interface{}) (ctx Context, cancel CancelFunc)
	// context.WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	// context.WithDeadline(parent Context, d time.Time) (ctx Context, cancel CancelFunc)
	// context.WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel CancelFunc)

	//
	//   Best practices
	//

	// - context.Background should be used only at the highest level, as the root of all derived
	//   contexts
	// - context.TODO should be used where not sure what to use or if the current function will
	//	 be updated to use context in future
	// - context cancelations are advisory, the functions may take time to clean up and exit
	// - context.Value should be used very rarely, it should never be used to pass in optional
	//   parameters. This makes the API implicit and can introduce bugs. Instead, such values
	//   should be passed in as arguments.
	// - Don’t store contexts in a struct, pass them explicitly in functions, preferably, as the
	//   first argument.
	// - Never pass nil context, instead, use a TODO if you are not sure what to use.
	// - The Context struct does not have a cancel method because only the function that derives the
	//   context should cancel it.

	printHeader("Context")

	// In the following example, you can see a function accepting context starts a goroutine and
	// waits for that goroutine to return or that context to cancel. The select statement helps us
	// to pick whatever case happens first and return.

	// <-ctx.Done() once the Done channel is closed, the case <-ctx.Done(): is selected.

	// Once this happens, the function should abandon work and prepare to return.
	// That means you should close any open pipes, free resources and return form the function.
	// There are cases when freeing up resources can hold up the return, like doing some clean up
	// that hangs, etc. You should look out for any such possibilities while handling the context return.

	//Make a background context
	ctx := context.Background()
	//Derive a context with cancel
	// returns the new context (with cancel) and the cancellation token
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	//defer canceling so that all the resources are freed up
	//For this and the derived contexts
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunction()
	}()

	//Cancel context after a random time
	//This cancels the request after a random timeout
	//If this happens, all the contexts derived from this should return
	go func() {
		sleepRandom("Main", nil)
		cancelFunction()
		fmt.Println("Main Sleep complete. canceling context")
	}()
	//Do work (until the process ends or sleepRandom triggers the WithTimeout )
	doWorkContext(ctxWithCancel)
}
