package main

import "fmt"

// MaximunInt Declare 'Exported' const variable (public)
const MaximunInt = 2147483647

// Declare 'Non-exported' const variable (private)
const minimumInt = -2147483647

// Lowercase and UpperCase variables with ':=' are not allowed outside a body (method)
// globalVariable := 1
// GlobalVariable := 2

// Lowercase and UpperCase variables with 'var' are allowed outside a body (method)

var globalVariable = 1001

// GlobalVariable Exported (Public) variables, methods, structs etc.. must have a comment
var GlobalVariable = 2001

// Person Basic struct (exported)
type Person struct {
	// Embedded from other type (similar to inherit)
	string // uuid
	// contains filtered or non-exported fields
	Name string
	Age  int
}

// Child inherits from struct Person (exported)
type Child struct {
	Person
	Parent string
}

// Alias from float32 (type inherited, no methods inherited)
type celsius float32

// Override the 'toString()' function for 'celsius' type
func (c celsius) String() string {
	return fmt.Sprintf("%.2f°", c)
}

// Alias from celsius (data inherited, no methods inherited)
type celsius2 celsius

// Equality types (data inherited, methods inherited)
type celsius3 = celsius

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	printHeader("Variable Declaration")

	// Non used variables, means a compilation error
	//var error = "Un-comment this line to get an error from the Go compiler"

	// `var` declares 1 or more variables (type is inferred).
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	// (type can be also inferred from value)
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Use var () to declare mutliple variables and types (bulk)
	var (
		x, z   int
		value1 = "value1"
		value2 = "value2"
	)
	fmt.Println(x, z)
	fmt.Println("value1:", value1, ",value2:", value2)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)
	// Similar, zero-valued for a `string` is 'empty'
	var chars string
	fmt.Println(chars)
	fmt.Println(len(chars))

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "short"` in this case.
	f := "short"
	fmt.Println(f)

	printHeader("Constants and Private/Public (Exported/Non-exported)")

	// Exported constants can be accessed from other modules
	fmt.Println(MaximunInt)

	// Non-Exported constants cannot be accessed from other modules
	// but internally it is possible.
	fmt.Println(minimumInt)

	// Exported variables can be accessed from other modules
	fmt.Println(GlobalVariable)

	// Non-Exported variables cannot be accessed from other modules
	fmt.Println(globalVariable)

	printHeader("Structures basics and inheritance (kind-of)")

	// Exported structures can be accessed from other modules
	// The embedded type van be passed through constructor
	parent := Person{
		string: "6651dbeb-7a59-49b9-9771-e27043cb0e56",
		Name:   "Manuel García",
		Age:    32,
	}
	// Print default structure stdout
	fmt.Println(parent)

	// Create a child from previous struct (specialization)
	child := Child{
		Person: Person{
			string: "e540d544-61d3-4267-baf5-4b68df859c9b",
			Name:   "Aitor García",
			Age:    12,
		},
		Parent: parent.string,
	}
	// Print default structure stdout
	fmt.Println(child)

	printHeader("Aliases and Casting (types)")

	// Types can be used to declare new custom types (Aliases)
	// type celsius float32
	degree := celsius(10.0)
	fmt.Println(degree)
	fmt.Printf("%T\n", degree)

	// Aliases can be casted to inherit types
	// however methods are not inherited
	degree2 := celsius2(20.0)
	fmt.Println(degree2)
	// The override method String() is not used for the alias
	fmt.Printf("%T\n", degree2)

	// However, it can be casted celsius2 to celsius, since both are aliases
	var degreeFromcelsius2 = celsius(degree2)
	fmt.Println(degreeFromcelsius2)

	// Equality types shares
	degree3 := celsius3(30.0)
	fmt.Println(degree3)
	fmt.Printf("%T\n", degree3)

}
