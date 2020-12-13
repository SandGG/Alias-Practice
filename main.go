package main

import "fmt"

func main() {
	type (
		myString1 string
		myString2 string

		myInt1 int
		myInt2 myInt1
	)

	var (
		ban   bool
		text1 myString1 = "Hola"
		text2 myString1 = "Hola"

		text3          myString2 = "Hola"
		miCadenaString           = string(text1)
		miCadena                 = "Hola"

		number1 myInt1 = 123
		number2 myInt2 = 123
	)

	//	Diferent types
	if text1 == text3 {
	}
	if number1 == number2 {
	}

	if text1 == text2 {
		ban = true
	}
	fmt.Println(ban)

	if miCadenaString == miCadena {
		ban = true
	}
	fmt.Println(ban)

}
