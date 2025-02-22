package main

import "fmt"

type Persons struct {
	name string
	Id   int64
	age  int
}

type Employee struct {
	code      string
	TIMESTAMP int64
	Persons
}

type Maneger struct {
	code    string
	company string
	Persons
}

func main() {

	e := Employee{
		code:      "12",
		TIMESTAMP: 10,
		Persons: Persons{
			name: "John",
			Id:   1,
			age:  25,
		},
	}

	M := Maneger{
		code:    "1234",
		company: "cocompany",
		Persons: Persons{
			name: "max",
			Id:   2,
			age:  45,
		},
	}
	fmt.Println(e)
	fmt.Println(M)
}

// ارث بَری یکی از اصول برنامه نویسی شی گراست که در آن
// یک class
// می تواند ویژگی های
// یک class
// دیگر را به ارث ببرد
