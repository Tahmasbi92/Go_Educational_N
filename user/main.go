package main

import (
	"Person/user"
	"fmt"
)

func main() {
	new := user.New("alireza", 10)

	reult := new.Start()
	fmt.Println(reult)
}
