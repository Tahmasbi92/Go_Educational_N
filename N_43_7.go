package main

import "fmt"

func main() {

	x := 10

	defer fmt.Println("this is with defer", x)

	x = 20

	fmt.Println("this is without defer", x)

}
