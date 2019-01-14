package main

import (
	"errors"
	"fmt"
	"time"
)

// Private that divides two numbers
// it checks if the divisor is not zero
func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisor cannot be 0")
	}
	return a / b, nil
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	printHeader("Conditonals")

	// Check all the operators
	// https://yourbasic.org/golang/operators/

	// Use 'if statement {}' to check a statement condition.
	// There is no parentesis in the statement

	// Valid 'logical operators' are : && (and), || (or) and ! (negation)
	// Valid 'expresions' are: == (equal to), != (not equal to), < (less than)
	//                         <= (less than or equal to) , > (greater than) ,
	//                         >= (greater than or equal to)

	a := 0
	b := 3

	// Simple if with one statement
	if b > a {
		fmt.Println("a is greater than b")
	}

	// Simple if/else with multiple statement
	// Use parentesis to
	if (a < b && b == 2) || (a == 0) {
		fmt.Println("the statement is satisfied")
	} else {
		fmt.Println("the statement is not satisfied")
	}

	// Multiple if/else if/else statements
	if a != 0 {
		fmt.Println("the  first statement is satisfied")
	} else if b >= 3 {
		fmt.Println("the second statement is satisfied")
	} else {
		fmt.Println("the last statement is satisfied")
	}

	// Assignment and conditional statement
	if num := 9; num > 0 {
		fmt.Println("num is greater than 0")
		fmt.Println(num)
	}
	// num is defined inside the if-statement scope
	//fmt.Println(num)

	printHeader("Switch")

	//Basic switch with default
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Today is Saturday.")
	case time.Sunday:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Today is a weekday.")
	}

	//No condition
	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	//Case list (multiple choices per case)

	c := ' '
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		fmt.Println("first case satisfied")
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("second case satisfied")
	default:
		fmt.Println("Good evening!")
	}

	//Fallthrough
	// A fallthrough statement transfers control to the next case.
	// It may be used only as the final statement in a clause.
	switch 2 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	}

	//   Exit with break
Loop:
	for _, ch := range "a b\nc" {
		switch ch {
		case ' ': // skip space
			break
		case '\n': // break at newline
			break Loop
		default:
			fmt.Printf("%c\n", ch)
		}
	}

	printHeader("Loops")

	//Three-component loop

	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)

	// While loop

	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	// Infinite loop with break condition

	total := 0
	for {
		if total > 10 {
			break
		}
		total++ // repeated forever
	}
	fmt.Println(total) // never reached

	// For-each range loop

	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// Range using channels (queues)

	fmt.Println("Create the channel")
	ch := make(chan int)
	go func() {
		// Defer to close the channel at the end
		defer close(ch) // close the channel
		fmt.Println("Send number 1 to the channel")
		ch <- 1
		fmt.Println("Send number 2 to the channel")
		ch <- 2
		fmt.Println("Send number 3 to the channel")
		ch <- 3
	}()
	fmt.Println("Print the values (dequeue)")
	// Start printing the values from the channel (async)
	for n := range ch {
		fmt.Println(n)
	}

	// Skip values inside a loop

	total = 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		total += i
	}
	fmt.Println(total) // 6 (2+4)

	printHeader("Error handling")

	// Normal division
	division, err := div(8, 4)
	if err != nil {
		fmt.Println("Error in the division")
	} else {
		fmt.Println(division)
	}

	// Zero division (forze the errro)
	division, err = div(8, 0)
	if err != nil {
		fmt.Println("Error in the division")
	} else {
		fmt.Println(division)
	}

	// Ifnore the second parameter
	division, _ = div(8, 2)
	fmt.Println(division)

	printHeader("Type assertions (Cast) and type switches")

	// A type 'assertion' doesn’t really convert an interface to another data type,
	// but it provides access to an interface’s concrete value, which is typically what you want.

	// The type assertion x.(T) asserts that the concrete value stored in x is of type T, and that x is not nil.
	//  - If T is not an interface, it asserts that the dynamic type of x is identical to T.
	//  - If T is an interface, it asserts that the dynamic type of x implements T.

	// Declare a variable foo using interface{}
	var x interface{} = "foo"
	// It detects x is string
	fmt.Printf("%T\n", x)

	// No need to use 'var s string = x.(string)'
	// since it infers x is string
	var s = x.(string)
	fmt.Println(s) // "foo"

	// It can be used a sencond output to check the result of the assetion
	s, ok := x.(string)
	fmt.Println(s, ok) // "foo true"

	p, ok := x.(int)
	fmt.Println(p, ok) // "0 false"

	// if must be used the second output'ok',
	// so go does not throw the 'panic' command
	//q := x.(int)   // ILLEGAL

	// Type switches
	var t interface{} = "foo"

	switch v := t.(type) {
	case nil:
		fmt.Println("t is nil") // here v has type interface{}
	case int:
		fmt.Println("t is", v) // here v has type int
	case bool, string:
		fmt.Println("t is bool or string") // here v has type interface{}
	default:
		fmt.Println("type unknown") // here v has type interface{}
	}

	printHeader("Defer Ordering")

	// The defer order is basically a stack (LIFO)
	// Last in, first out (Stack)
	fmt.Println("Hello")
	defer fmt.Println("World")
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}

	// It prints from last to first (order)
	// Hello, 3, 2, 1, World

}
