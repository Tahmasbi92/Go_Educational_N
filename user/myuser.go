package user

import "fmt"

type user struct {
	name string
	age  int
}

func New(name string, age int) *user {
	return &Users{name, age}
}
func (s *Users) start() string {
	return fmt.Sprint("Starting")
}
