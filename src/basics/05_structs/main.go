package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// It must be exported all the structs, fields and functions (capitalized) to be used externally
// Every exported (capitalized) name in a program should have a doc comment.

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
	Age   int32  `json:"age"`
	Roles []Role `json:"roles"`
}

// App struct
type App struct {
	ID    string `json:"id"`
	Title string `json:"title"`
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

func main() {

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
