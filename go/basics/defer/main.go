package main

import "fmt"

func a() {
	i := 0
	defer fmt.Println("x")
	defer fmt.Println(i)
	i++
	return
}

func divide(a, b int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover:", r)
		}
	}()
	if b == 0 {
		panic("Division by 0")
	}
	return a / b, nil
}

func main() {
	a()

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
