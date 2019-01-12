package main

import "fmt"

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	printHeader("Array")

	// Arrays are created with a fixed length
	var numbers [4]int32
	// Print the length using the 'len' function
	fmt.Println(len(numbers))
	// Print empty values
	fmt.Println(numbers)
	// Assign a value for each fixed position in the array
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	// Print all values of numbers
	fmt.Println(numbers)

	// Arrays can be created inline (also fixed)
	var inlineArray = [4]int32{1, 2, 3, 4}
	fmt.Println(len(inlineArray))
	fmt.Println(inlineArray)

	// Following line will give an error since the initial length (5) is greater than 4
	// var error = [4]int32{1, 2, 3, 4, 5} // Remove one value to fix it

	printHeader("Slides")

	// Slides are created with a undefined length
	var letters []string
	// Print the length using the 'len' function
	fmt.Println(len(letters))
	// Print empty values
	fmt.Println(letters)

	// Following assignment will give you a runtime error (not at compile time)
	// Runtime Error, since the length is 0. Use 'append(slide, value)' function instead
	// letters[0] = "First"

	// Append the values at the end of the slide
	// NTE:  Append is idempotent so it returns a new slide with the new item appended
	letters = append(letters, "First")
	letters = append(letters, "Second")
	letters = append(letters, "Third")
	letters = append(letters, "Fourth")

	// Print all values of numbers
	fmt.Println(len(letters))
	fmt.Println(letters)

	// Arrays can be created inline (also fixed)
	var inlineslide = []string{"First", "Second", "Third", "Fourth"}
	fmt.Println(len(inlineslide))
	fmt.Println(inlineslide)

	printHeader("Fancy indexing")

	// Arrays can be created inline (also fixed)
	var fancyIndexing = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(len(fancyIndexing))
	fmt.Println(fancyIndexing)

	// Print a single value via index
	fmt.Println(fancyIndexing[3])
	fmt.Println(fancyIndexing[10])

	// Print a range of values via fancy indexing
	fmt.Println(fancyIndexing[2:6])
	fmt.Println(fancyIndexing[2:6])
	fmt.Println(fancyIndexing[:6])
	fmt.Println(fancyIndexing[6:])

	// There is no reverse order using [::-1] like in Python.

	printHeader("Maps")

	// Create a map (key (string), value (int))
	people := make(map[string]int)
	fmt.Println(len(people))
	fmt.Println(people)

	// Star adding items into the map
	people["Peter"] = 23
	people["Maria"] = 12
	people["Carlos"] = 101

	fmt.Println(len(people))
	fmt.Println(people)

	// Get an item via its key
	fmt.Printf("Carlos: %d\n", people["Carlos"])
	fmt.Printf("Maria: %d\n", people["Maria"])

	// Create a map (key (string), value (int)) inline
	inlineMap := map[string]int{
		"Peter":  23,
		"Maria":  12,
		"Carlos": 101, // Comma needed in multi-line definition
	}
	fmt.Println(len(inlineMap))
	fmt.Println(inlineMap)

	printHeader("Loops (for statement)")

	// It is recommended to use single letters for loops indexes: i, j, k, ..

	// Create a collection
	var loopOver = []int{1, 2, 3, 4, 5}
	fmt.Println(loopOver)

	// For loop
	for i := 0; i < len(loopOver); i++ {
		fmt.Println(loopOver[i])
	}

	// While loop
	j := 0
	for j < len(loopOver) {
		fmt.Println(loopOver[j])
		j++
	}

	// There is no map, filter, reduce, or any functional operation in Go
	// The idea is to simplify and make more efficient the language using standard loops...

	printHeader("Ranges")

	// Create a collection
	var list = []int{1, 2, 3, 4, 5}
	fmt.Println(list)

	// For each
	for index, value := range list {
		fmt.Printf("[%d]=%d\n", index, value)
	}

	// For each (without index)
	for _, value := range list {
		fmt.Print(value)
	}
	fmt.Println()

	colors := map[int]string{1: "Red", 2: "Orange", 3: "Blue", 4: "Yellow"}
	// Loop over a map
	for key, value := range colors {
		fmt.Printf("[%d]=%s\n", key, value)
	}
	fmt.Println()

}
