// Name of the package
// - Multiple go files per folder with the same package name
// - Use same package name per folder, since imports use folders no package names
package main

// Impot packages:
// - Multi-import must be in different lines ( "fmt"\n "errors")
// - it can be used aliases: formatter "fmt"
// - Packages must be relative to $GOPATH/src: 'github.com/01_helloworld/package'
import "fmt"

// Main entry point (commmand)
func main() {
	// Print text using the standard package 'fmt'
	fmt.Printf("hello, World\n")
}
