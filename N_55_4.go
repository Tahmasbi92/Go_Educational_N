package main

import "fmt"

func Print[T string | int | float64](value T) T { // be jaye estefade az "any" ke shamele
	// tamame vorudiha (string , int , float , bool) mishavad mitavan az in dasture bala ba tavajoh be
	//func main payeen estafade kard.
	fmt.Println(value)

	return value
}

func main() {
	print(12)
	print("Hi my name is Alireza")
	print(12.12)

}
