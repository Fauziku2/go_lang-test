package main

import (
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  10,
	}

	fmt.Println(fido)
}
