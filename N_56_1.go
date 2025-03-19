package main

import (
	"errors"
	"fmt"
)

func main() {

	result, error := createError(true)

	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(result)
}

func createError(er bool) (string, error) {

	if er {
		return "", errors.New("error creating")
	}

	return "you dont like create error", nil
}
