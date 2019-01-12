package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Public/Private = Exported/Un-exported
// It must be 'exported' all the structs, fields and functions (capitalized) to be used externally
// Every exported (capitalized) name in a program should have a doc comment.

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

// Role struct
type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User structure
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int32  `json:"age,omitempty"`
	Roles []Role `json:"roles,omitempty"`
}

// App struct
type App struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// SerializeUser function
func SerializeUser(data []byte) (User, error) {
	// 'var result user' or 'result := user{}'
	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Error serializing bytes into User")
		return user, errors.New("Error serializing bytes into User")
	}
	fmt.Println("User Object serialized")
	return user, nil
}

// DeserializeUser function
func DeserializeUser(user User) ([]byte, error) {
	bytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error deserializing User into bytes")
		return bytes, errors.New("Error deserializing User into bytes")
	}
	fmt.Println("User Object deserialized")
	return bytes, nil
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

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

	printHeader("Add methods to Structs (Classes)")

	printHeader("Interfaces")

	printHeader("Structs Tags and Serialize/Deserialize (JSON)")

	// Create raw JSON to Serialize to a User struct
	data := []byte(`
      {
         "id": "1234",
         "name": "Javier",
         "email": "javier@gmail.com",
         "age": 35
      }
   `)

	// Serialize the byte into User Struct
	person, err := SerializeUser(data)
	if err != nil {
		return
	}
	fmt.Println(person)

	// Deserialize the byte into User Struct
	bytes, err := DeserializeUser(person)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
