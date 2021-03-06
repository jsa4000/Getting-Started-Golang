package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Conversion from nanoseconds to milliseconds
const nanoseconds = 1000000

// Greeting const used in the example (private)
const greeting = "Hello, World"

// PrintGreeting Command function (void).
// Function Public, visible to external modules
func PrintGreeting() {
	fmt.Println(greeting)
}

// PrintGreeting Command function (void).
// Function private, not visible to external modules
func printGreeting() {
	fmt.Println(greeting)
}

// PrintCount from 0 to the given number
// Use `defer` to do something when the function exists
func PrintCount(number int) {
	fmt.Println("Start Counting")
	defer fmt.Println("End Counting")
	for i := 0; i < number; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

// ForcePanic to quit the go process
// Use `panic` when an error occurs and it must be exited
func ForcePanic(force bool) {
	fmt.Println("Start Counting")
	// Following defer sequence is applied normally on panic (before panic)
	defer fmt.Println("End Counting")
	if force {
		panic("Forced panic")
	}
	fmt.Println("This is never printed when panic is forced")
}

// IsGreeting Function that check if the text is a greeting (bool).
func IsGreeting(text string) bool {
	return text == greeting
}

// GetInteger check if the value is a number and return the value
// If it is not a number it returns an error
// _ context.Context parameter, is ignored
func GetInteger(_ context.Context, value interface{}) (int, error) {
	switch value.(type) {
	case int:
		return value.(int), nil
	case float64:
		return int(value.(float64)), nil
	case string:
		i, err := strconv.Atoi(value.(string))
		return i, err
	default:
		return 0, errors.New("Error. Unknown type")
	}
}

// User struct to define an user
type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

// String method to override the stdout of the struct
// In this case we use pointers for the methos so: 'user := $User{}' -> user.String()
func (u *User) String() string {
	return fmt.Sprintf("{Id:%s, Name:%s, Email:%s, Age:%d}", u.ID, u.Name, u.Email, u.Age)
}

// IsAdult returns if the user is greater than 18 years old.
// We can use  (u *User) or (u User)
func (u *User) IsAdult() bool {
	return u.Age >= 18
}

// SendEmail send an email to the user
func (u *User) SendEmail(text string) error {
	if u.Email == "" {
		return errors.New("The user has not configured the email Address")
	}
	fmt.Printf("Sent email to %s\n", u.Email)
	return nil
}

// HandlerFunc Definition handler used by the Logger function (decorator)
// This is a type fon functions (delegates).
type HandlerFunc func(args string)

// ProcessReq that explicitly uses 'HandlerFunc' type definition
func ProcessReq(args string) {
	// Simulate a delay to process the request (aysnc, database, external rpc calls, etc..)
	time.Sleep(100 * nanoseconds)
	fmt.Println(args)
}

// Logger is a decorator (wrapper) around another function (or closure)
// In golang is oftenly called Middleware, and it takes a pointer to a function
// (i.e. req := Logger(ProcessReq)) and it returns the same function wrapper into another.
// Finally the function can be executed normally using the type paremeters req("args")
//
// NOTE: Basically, every function with the form 'func(args string)' or 'HandlerFunc' type,
//       can be used inside this logger or 'decorator'
func Logger(f HandlerFunc) HandlerFunc {
	fmt.Println("Creating Logger Handler")
	defer fmt.Println("Created (returned) Logger Handler")

	// Return the same function passed in parameters but
	return func(args string) {
		start := time.Now()
		fmt.Println("Start Function Handler")
		// Following line must be enclosed by a lambda function
		// fmt.Printf("Ends Function Handler in %.2f ms\n", float64(time.Since(start))/nanoseconds) // DOES NOT WORK
		defer func() { fmt.Printf("Ends Function Handler in %.2f ms\n", float64(time.Since(start))/nanoseconds) }()
		f(args)
	}
}

// Counter is a closure function that holds a variable (state). So in every call the inner status changes
func Counter() func(int) int {
	// Initialize state.
	// Initial state can be passed through parameters if wanted
	// i.e 'func Counter(start int) int' -> total := start (or simply use 'start' as total counter)
	total := 0
	return func(value int) int {
		// Modify current state with current value
		total += value
		// Return current state after the change
		return total
	}
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	printHeader("Functions")

	// Print the greeting (command)
	PrintGreeting() // public

	// Print the greeting (command)
	printGreeting() // private

	// Call a function to check if the text is a greeting (true)
	if IsGreeting("Hello, World") {
		println("Is a greeting")
	}

	// Call a function to check if the text is a greeting (false)
	if !IsGreeting("Goodbye, World") {
		println("Is not a greeting")
	}

	// Functions can have multiple return values (tuples)
	// First value 'context' is omitted by using '_', so a 'nil' value is passed

	// Pass int to the function (returns the int)
	number, err := GetInteger(nil, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number)
	}

	// Pass float to the function (converts float64 to int)
	number, err = GetInteger(nil, 3.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number)
	}

	// Pass string function to the function (converts string to int)
	number, err = GetInteger(nil, "3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number)
	}

	// Pass struct variable to the function (returns an error)
	number, err = GetInteger(nil, User{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(number)
	}

	printHeader("Defer, Panic, recover statements")

	// Use `defer` to do something when the function exists
	PrintCount(5)

	// Use `panic` when an error occurs and it must be exited
	//ForcePanic(true) // Un-comment this line to see the panic exception

	// Do not Use `panic`, so it continued with the process
	ForcePanic(false)

	printHeader("Lambda Functions")

	// Lambda function or anonymous functions are functions that are declared inside
	// the function scope, instead declaring it globally

	// A language that considers procedures to be "first-class" allows
	// functions to be passed around just like any other value.

	// Create a lambda function sum (functions as first-class citizens)
	sum := func(a, b int) int {
		return a + b
	}
	//Execute previous lambda function
	total := sum(2, 5)
	fmt.Printf("Total Sum: %d\n", total)

	// Create and execute a lambda function
	diff := func(a, b int) int { return a - b }(7, 2)
	fmt.Printf("Total Diff: %d\n", diff)

	printHeader("Methods")

	// Methods are functions, however they are attached to an struct (like classes)
	// - data       : structs
	// - behavior   : methods (functions)

	user := &User{
		ID:    "1234",
		Name:  "Javier",
		Email: "javier@gmail.com",
		Age:   35,
	}
	fmt.Println(user)
	// Same output, since we have overridden String() method.
	fmt.Println(user.String())
	// Return if the user is adult (age>18)
	fmt.Printf("IsAdult: %t\n", user.IsAdult())

	if err := user.SendEmail("This is a content"); err != nil {
		fmt.Println(err)
	}

	// Modify the current content of the struct
	// Note: Structs must be immutable
	user.Email = ""
	if err := user.SendEmail("This is a content"); err != nil {
		fmt.Println(err)
	}

	printHeader("Closures")

	// Go functions may be 'closures'. A closure is a function value that references variables
	// from outside its body. The function may access and assign to the referenced variables;
	// in this sense the function is "bound" to the variables.

	// Logger is a decorator (wrapper) around another function (this is called 'closure')
	// In golang is oftenly called Middleware, and it takes a pointer to a function (typed)
	// (i.e. req := Logger(ProcessReq)), and it returns the same function wrapped into another.
	request := Logger(ProcessReq) // ProcessReq is the pointer to the function to call later on, inside the closure

	// Finally the function can be executed normally using the 'HandlerFunc' type
	request("This is a request")

	// Creates a Counter closure to store a stateful counter
	counter := Counter()
	fmt.Println(counter(0))
	fmt.Println(counter(2))
	fmt.Println(counter(5))

}
