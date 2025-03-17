package main

import "fmt"

func Print[T string | int | float64](value T) T {
	fmt.Println(value)

	return value
}

func main() {
	print(12)
	print("Hi my name is Alireza")
	print(12.12)

}
