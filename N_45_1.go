package main

import "fmt"

type car struct {
	name       string
	createDate string
	company    string
}

func main() {

	car1 := car{
		name:       "Benz",
		createDate: "2024",
		company:    "Benz",
	}

	fmt.Println(car1)

}
