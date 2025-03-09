package main

import (
	"fmt"
	structs "project/Structs"
)

func main() {
	u := structs.AllUsers{ // به فانکشن یوزر می خواهم دسترسی داشته باشم
		Name: "Alireza",
		Age:  27,
	}
	result := Functions.DispalyInfo(u.Name) // می خواهم به فانکشن دیسپلی دسترسی داشته باشم

	fmt.Println(result)
}
