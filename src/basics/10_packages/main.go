package main

import (
	"context"
	"fmt"

	"github.com/basics/09_packages/users"
)

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	// This is the old way to create modules.
	// Modules must be inside the $GOPATH/src

	// Go versions > 11.x supports modules (GO111MODULE=auto)
	// go mod init

	// Pre-requisites
	// go get github.com/satori/go.uuid
	// go get golang.org/x/crypto/bcrypt

	// Create the layered (hexagon) arquitecture

	printHeader("Starting application")
	defer printHeader("Finished Application")
	defer printHeader("Finishing Application")

	// Create Repository
	usersRepository := users.NewMockRepository()

	// Create Service (Inject the repository)
	UsersService := users.NewServiceImpl(usersRepository)

	// Create a context, just to pass through the chain
	ctx := context.Background()

	printHeader("Get all the users")

	// Get all the Users
	allUsers, _ := UsersService.GetAll(ctx)
	for _, user := range allUsers {
		fmt.Println(user)
	}

	printHeader("Create a new user")

	// Craete new User
	user := users.User{
		Name:     "Alvaro",
		Email:    "alvaro.golang@example.com",
		Password: "myPassword",
	}
	user, _ = UsersService.Create(ctx, user)
	fmt.Println(user)

	printHeader("Get all the Users")

	// Get all the users
	allUsers, _ = UsersService.GetAll(ctx)
	for _, user := range allUsers {
		fmt.Println(user)
	}

	printHeader("Get User by Id")

	// Get existing user the users
	user, err := UsersService.GetByID(ctx, user.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}

	// Get all the users
	user, err = UsersService.GetByID(ctx, "id-1234")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}

	printHeader("Remove User by Id")

	// Remove user by Id

	user = allUsers[0]
	err = UsersService.RemoveByID(ctx, user.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("User with id %s removed", user.ID)
	}

	printHeader("Get all the Users")

	// Get all the users
	allUsers, _ = UsersService.GetAll(ctx)
	for _, user := range allUsers {
		fmt.Println(user)
	}

}
