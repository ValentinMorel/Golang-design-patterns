package main

import (
	"errors"
	"fmt"
)

/*
Factory design pattern aims to abstract the user from the knowledge of the struct
he needs to achieve for a specific purpose, such as retrieving some value.
User only needs an interface that provides this value.
It also eases the process of downgrading or upgrading of the implementation of the
underlying type if needed
*/

const (
	french  = 1
	english = 2
)

type translator interface {
	getMessage() string
}

type englishMessage string
type frenchMessage string

func (c *englishMessage) getMessage() string {
	return "Hello World"
}

func (d *frenchMessage) getMessage() string {
	return "Bonjour tout le monde"
}

func getTranslator(m int) (translator, error) {
	switch m {
	case french:
		return new(frenchMessage), nil
	case english:
		return new(englishMessage), nil
	default:
		return nil, errors.New("Unknown building.")
	}
}

func main() {

	obj1, _ := getTranslator(1)
	fmt.Printf("%T\n", obj1)
	fmt.Println(obj1.getMessage())

	obj2, _ := getTranslator(2)
	fmt.Printf("%T\n", obj2)
	fmt.Println(obj2.getMessage())

}
