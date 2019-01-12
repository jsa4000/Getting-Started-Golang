// Name of the package
// - Packages name must be short and must define the function of the package (DDD).
// - Packages names are no namespaces, avoid using generic names such as utils, helpers, etc...
// - Multiple go files per folder with the same package name (shared)
// - Use same package name per folder, since imports use folder names no package names
package main

// Import packages:
// - Multiple imports must be in different lines 'import (\n "fmt"\n "errors"\n)\n'
// - It can be used aliases for conflicts or conventions: formatter "fmt"
// - Packages must be relative to $GOPATH/src: 'github.com/01_helloworld/package'
// - From go > v.11.x packages can be in any folder using modules feature. 'get mod init'
import "fmt"

// Main entry point (command)
// - Upperletter case is used for public members
// - Lowerletter case is used for private members
// - Methods names must be short
// - Avoid repetition in package names and methods names: user.NewUser() -> user->New()
func main() {
	// Print text (stdout) using the standard package 'fmt'
	fmt.Printf("hello, World\n")
}
