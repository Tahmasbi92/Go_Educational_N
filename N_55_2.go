package main

import "fmt"

func main() {

	peintstring("hello world")
	printinteger(12)
}

func peintstring(s string) string {

	fmt.Println(s)

	return s
}

func printinteger(i int) int {

	fmt.Println(i + 100)
	return i
}
