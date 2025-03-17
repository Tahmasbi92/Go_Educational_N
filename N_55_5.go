package main

import "fmt"

type Number interface {
	int | string
}

func Add[T Number](a, b T) T {

	return b + a
}

func main() {

	result :=
		Add(2, 3)

	fmt.Println(result)

	result2 := Add("Hello", "world")

	fmt.Println(result2)
}
