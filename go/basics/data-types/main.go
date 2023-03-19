package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Boolean
	var b bool = true
	fmt.Println("Boolean: ", b)

	// Integer
	var i int = 42
	fmt.Println("Integer: ", i)

	// Floating point
	var f float64 = 3.14
	fmt.Println("Float: ", f)

	// Complex numbers
	var c complex128 = complex(1, 2)
	fmt.Println("Complex: ", c)

	// Strings
	var s string = "Hello, World!"
	fmt.Println("String: ", s)

	// Arrays
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println("Array: ", a)

	// Slices
	var sl []int = []int{1, 2, 3}
	fmt.Println("Slice: ", sl)

	// Maps
	var m map[string]int = map[string]int{"a": 1, "b": 2}
	fmt.Println("Map: ", m)

	// Structs
	type person struct {
		name string
		age  int
	}
	var p person = person{name: "John", age: 30}
	fmt.Println("Struct: ", p)

	// Pointers
	var ptr *int = &i
	fmt.Println("Pointer: ", ptr)

	// Interfaces
	var intf interface{} = 42
	fmt.Println("Interface: ", intf)

	// Type assertion
	var val interface{} = "hello"
	str, ok := val.(string)
	fmt.Println("Type assertion: ", str, ok)

	// Type reflection
	var x float64 = 3.14
	fmt.Println("Type reflection: ", reflect.TypeOf(x))
}
