package main

import "fmt"

type car struct {
	name       string
	createDate string
	company    string
}

func main() {

	car5 := car{"pride", "1999", "sipa"}
	fmt.Println(car5)
}
