// Declare the name of the package (module)
package main

// Impot the packages,
// - Multi-import must be in different lines ( "fmt"\n "errors")
// - it can be used aliases: out "fmt"
// - Packages must be relative from $GOPATH/src: 'github.com/01_helloworld/package'
import (
	"fmt"

	helper "github.com/01_helloworld/package"
)

// Create the main function (if aoo)
func main() {
	// Implement the main function

	// Print the hello world using the standard package fmt
	fmt.Printf("hello, World\n")

	// Print the hello world using the helpers package created
	helper.PrintHelloWord("Javier")
}
