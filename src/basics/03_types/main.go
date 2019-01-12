package main

import "fmt"

// Like pointer receiver
func pointerReceiver(value *int) {
	*value += 2
}

// Like value receiver
func valueReceiver(value int) {
	value += 2
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	printHeader("String Types")

	// Go provides two ways to declaring string types (inline)

	// UTF8: Use standard double quotes ""
	unicode := "It is 42\u00B0 F outside!"
	fmt.Println(unicode)

	// ASCII: Use standard single quotes ``
	// Value \u00B0 is not recognized as unicode character anymore
	multiline := `
	It is 42\u00B0 F 
	outside!
	`
	fmt.Println(multiline)

	line := `It is 42\u00B0 F outside!`
	fmt.Println(line)

	printHeader("Integer Types")

	// Go has several types, with different internal sizes, that can be used to store integer values as listed below:

	// int{8,16,32,64} — singed integers of 8,16,32,64 bit in size (int32, int64, etc)
	// uint{8,16,32,64} — unsigned integers of 8,16,32,64 bit in size (i.e. uint8)
	// byte — alias for and equivalent to uint8
	// rune — used to represent characters, alias for and equivalent to int32
	// int — signed integers of at least 32-bit in size, not equivalent to int32
	// uint — unsigned integers of at least 32-bit; not equivalent to uint32
	// uintptr — dedicated for storing memory address pointers
	// float{32,64} — point floating numbers 32 and 64 bit in size (i.e. float32)
	// complex{64,128} — complex numbers 64 and 128 bit in size (i.e. complex64)

	// Use specific type during it declaration
	var lInt32 int32 = -1234
	var lByte uint8 = 3
	var lUint32 uint32 = 1234

	fmt.Printf("%T\n", lInt32)
	fmt.Printf("%T\n", lByte)
	fmt.Printf("%T\n", lUint32)

	// Uninitialized, integral types have a zero value of 0.
	// Integers can be initialized with constant literals which can be expressed as a decimal, octal, and hex as shown below
	var color uint32 = 0xFEFEFE // hex (0x prefix)
	var mod = 0466              // octal (0 prefix)
	count := 1245               // decimal

	fmt.Printf("%T\n", color)
	fmt.Printf("%T\n", mod)
	fmt.Printf("%T\n", count)

	// Casting object (floating-point number -> integer)
	var lFloat64 = 1.1
	var lInt64 = int64(lFloat64)

	fmt.Println(lFloat64)
	fmt.Printf("%T\n", lFloat64)
	fmt.Println(lInt64)
	fmt.Printf("%T\n", lInt64)

	printHeader("Boolean Types")

	var error = true
	fmt.Println(error)
	fmt.Printf("%T\n", error)

	printHeader("Pointer Types")

	var pointer *int
	var number = 34
	pointer = &number

	fmt.Printf("%T\n", pointer)
	fmt.Printf("%T\n", number)

	/// Write the address of the pointer
	fmt.Println(&pointer)
	/// Write the address of the content (pointer address <> content address)
	fmt.Println(pointer)
	/// Write the content assigned to that address
	fmt.Println(*pointer)

	// Pointer receiver vs Value receiver

	// Pointer Receiver (Passed By Value)
	value1 := 2
	valueReceiver(value1)
	fmt.Printf("valueReceiver = %d\n", value1)

	// Pointer Receiver (Passed By Reference)
	value2 := 2
	pointerReceiver(&value2)
	fmt.Printf("pointerReceiver = %d\n", value2)

}
