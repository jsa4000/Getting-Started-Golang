package main

import "fmt"

// Declare 'Exported' const variable (public)
const MaximunInt = 2147483647

// Declare 'Non-exported' const variable (private)
const minumunInt = -2147483647

// Basic struct (exported)
type Person struct {
	// contains filtered or unexported fields
	Name string
	Age int
}

// Aliases (data exported, no methods exported)
type celsius float32

// Override the toString function for celsius type
func (c celsius) String() string {
	return fmt.Sprintf("%.2f°", c)
}

// Aliases (data exported, no methods exported)
type celsius2 celsius

// Equal (data)
type celsius3 = celsius

func main() {

	// `var` declares 1 or more variables (type is infered).
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)
	// empy value for `string` 
	var chars string
	fmt.Println(chars)
	fmt.Println(len(chars))

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "short"` in this case.
	f := "short"
	fmt.Println(f)

	// Eported consts can be accesed from other modules
	fmt.Println(MaximunInt)

	// Non-Eported consts cannot be accesed from other modules
	// but internally it is possible.
	fmt.Println(minumunInt)
	
	// Eported structs can be accesed from other modules
	person := Person {
		Name: "Manuel Robledo",
		Age: 32,
	}
	// Print default strunct stdout
	fmt.Println(person)

	// Types can be used to delacle new custom types (Aliases)
	// type celsius float32
	degree := celsius(10.0)
	fmt.Println(degree)
	fmt.Printf("%T\n",degree)

	// Aliases can be casted to inherited types
	// however methods are not inherited
	degree2 := celsius2(20.0)
	fmt.Println(degree2)
	// the override method String() is not used for the alias
	fmt.Printf("%T\n",degree2)
	// Cast celsius2 to celsiu
	degree = celsius(degree2)
	fmt.Println(degree)

	// Equality types shares 
	degree3 := celsius3(30.0)
	fmt.Println(degree3)
	fmt.Printf("%T\n",degree3)

}
