package main

import "fmt"

// Like pointer receiver
func pointerRecevicer(value *int) {
	*value += 2
}

// Like value receiver
func valueReceiver(value int) {
	value += 2
}

func main() {

	// STRING //

	// Go provides two ways to declaring string types (inline)

	// UTF8
	message := "It is 42\u00B0 F outside!"
	fmt.Println(message)

	// ASCII
	multiline := `
	It is 42 \u00B0 F 
	outside!
	`
	fmt.Println(multiline)

	// INTEGER //

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
	var lbyte uint8 = 3
	var lUint32 uint32 = 1234

	fmt.Printf("%T\n", lInt32)
	fmt.Printf("%T\n", lbyte)
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
	var n = 1.1
	var m = int64(n)
	fmt.Println(m)

	// BOOLEAN //

	var myerror = true
	fmt.Println(myerror)
	fmt.Printf("%T\n", myerror)

	// POINTER

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

	// Pointer Receicer (Passed By Value)
	value1 := 2
	valueReceiver(value1)
	fmt.Printf("valueReceiver = %d\n", value1)

	// Pointer Receicer (Passed By Reference)
	value2 := 2
	pointerRecevicer(&value2)
	fmt.Printf("pointerRecevicer = %d\n", value2)

}
