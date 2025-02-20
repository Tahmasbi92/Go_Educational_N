package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("my err: ", err)
		return
	}

	fmt.Println("file opened: ", file)

}
