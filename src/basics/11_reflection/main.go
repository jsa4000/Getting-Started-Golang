package main

import (
	"fmt"
	"reflect"
)

// Role struct
type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User structure
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Age   int32
	Roles []Role `json:"roles,omitempty"`
}

func printHeader(text string) {
	fmt.Println()
	fmt.Println("*** " + text + " ***")
	fmt.Println()
}

func main() {

	printHeader("Reflection Basics")

	// Terminology for reflection

	// Kind: base type (finite): struct, ptr, int, func, slice, map, etc..
	// Type: base type (finite): struct, ptr, string, int, float64, func, slice, map, etc..

	a := "This is a string"
	k := 1
	e := 2.3
	s := User{}
	p := &User{}
	var q interface{} = User{} // Interface in this case is main.User (Type) and struct (kind)

	av := reflect.ValueOf(a)
	kv := reflect.ValueOf(k)
	ev := reflect.ValueOf(e)
	sv := reflect.ValueOf(s)
	pv := reflect.ValueOf(p)
	qv := reflect.ValueOf(q)

	fmt.Printf("String: Basic Type or Kind: %s, Direct or Custom Type: %s\n", av.Kind(), av.Type())
	fmt.Printf("Integer: Basic Type or Kind: %s, Direct or Custom Type: %s\n", kv.Kind(), kv.Type())
	fmt.Printf("Float: Basic Type or Kind: %s, Direct or Custom Type: %s\n", ev.Kind(), ev.Type())
	fmt.Printf("Struct: Basic Type or Kind: %s, Direct or Custom Type: %s\n", sv.Kind(), sv.Type())
	fmt.Printf("Pointer: Basic Type or Kind: %s, Direct or Custom Type: %s\n", pv.Kind(), pv.Type())
	fmt.Printf("Interface: Basic Type or Kind: %s, Direct or Custom Type: %s\n", qv.Kind(), qv.Type())

	printHeader("Reflection Structs")

	// Read the Struct using reflection
	user := User{}
	tp := reflect.TypeOf(user)
	val := reflect.ValueOf(user)
	for i := 0; i < tp.NumField(); i++ {
		ft := tp.Field(i)
		fv := val.Field(i)
		fmt.Printf("Name: %s  Kind: %s  Type: %s  Tag:  %s\n", ft.Name, fv.Kind(), fv.Type(), ft.Tag.Get("json"))
	}

	// Get the Field and Tag By Field Name 'Email'
	t := reflect.TypeOf(User{})

	fmt.Println("Field 'Email' with Tag")
	f, _ := t.FieldByName("Email")
	fmt.Println(f.Tag)
	v, ok := f.Tag.Lookup("json")
	if ok {
		fmt.Printf("%s, %t\n", v, ok)
	}

	fmt.Println("Field 'Age' without Tag")
	f, _ = t.FieldByName("Age")
	fmt.Println(f.Tag)
	v, ok = f.Tag.Lookup("json")
	if ok {
		fmt.Printf("%s, %t\n", v, ok)
	}

	printHeader("Reflection Switch Type")

	values := []interface{}{"This is a string", 1, 1.1, User{}}
	for _, value := range values {
		reflectValue := reflect.ValueOf(value)
		fmt.Printf("Basic Type or Kind: %s, Direct or Custom Type: %s\n", reflectValue.Kind(), reflectValue.Type())
		switch reflectValue.Kind() {
		case reflect.String:
			fmt.Printf("Value '%s' is String\n", value)
		case reflect.Int:
			fmt.Printf("Value '%s' is Int\n", value)
		case reflect.Float32, reflect.Float64:
			fmt.Printf("Value '%s' is Float\n", value)
		case reflect.Bool:
			fmt.Printf("Value '%s' is Bool\n", value)
		default:
			fmt.Printf("Kind is not supported for type %s for value '%s'\n", reflectValue.Kind(), reflectValue.String())
		}
	}

	printHeader("Reflection Encode")

	// Create
	a = "This is a string"
	av = reflect.ValueOf(a)
	fmt.Println(a)
	fmt.Printf("String: Basic Type or Kind: %s, Direct or Custom Type: %s\n", av.Kind(), av.Type())

	// Change its value, it cannot be setsince it is not unaddressable
	fmt.Println(av.CanSet())
	if av.CanSet() {
		av.SetString("New Value")
		fmt.Println(a)
	}

	// Use Elem() to access to the object stored in a address using reflection and pointers
	avp := reflect.ValueOf(&a)
	fmt.Printf("String: Basic Type or Kind: %s, Direct or Custom Type: %s\n", avp.Kind(), avp.Type())
	fmt.Printf("String: Basic Type or Kind: %s, Direct or Custom Type: %s\n", avp.Elem().Kind(), avp.Elem().Type())

	// Change its value, now it can be accessed, since we are using pointers references
	fmt.Println(avp.Elem().CanSet())
	if avp.Elem().CanSet() {
		avp.Elem().SetString("New Value")
		fmt.Println(a)
	}

	printHeader("Reflection Struct Encode")

	// Read the Struct using reflection
	user = User{
		ID:    "1234",
		Name:  "Javier Test",
		Email: "javier@example.com",
		Age:   22,
	}
	fmt.Println(user)

	tp = reflect.TypeOf(user)
	// Get the reference to the struct to modify by the reflection
	val = reflect.ValueOf(&user)
	for i := 0; i < tp.NumField(); i++ {
		ft := tp.Field(i)
		// Get the element from the pointer
		fv := val.Elem().Field(i)
		fmt.Printf("Name: %s  Kind: %s  Type: %s  Tag:  %s\n", ft.Name, fv.Kind(), fv.Type(), ft.Tag.Get("json"))
		switch fv.Kind() {
		case reflect.String:
			fmt.Printf("Value '%s' is String\n", fv.String())
			fv.SetString("New Value")
		case reflect.Int, reflect.Int32, reflect.Int64:
			fmt.Printf("Value '%s' is Int\n", fv.String())
			fv.SetInt(1)
		case reflect.Float32, reflect.Float64:
			fmt.Printf("Value '%s' is Float\n", fv.String())
		case reflect.Bool:
			fmt.Printf("Value '%s' is Bool\n", fv.String())
		default:
			fmt.Printf("Kind is not supported for type %s for value '%s'\n", fv.Kind(), fv.String())
		}
	}
	fmt.Println(user)
}
