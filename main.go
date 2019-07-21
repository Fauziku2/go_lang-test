package main

import (
	"fmt"

	"github.com/Fauziku2/go_lang-test/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}

	fmt.Println(fido)
}
