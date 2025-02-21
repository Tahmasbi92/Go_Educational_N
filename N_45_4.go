package main

import "fmt"

type car struct {
	name       string
	createDate string
	company    string
}

func main() {

	car4 := car{name: "PRIDE"}
	fmt.Println(car4)
}
