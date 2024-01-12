package main

import (
	"fmt"
)

func main() {
	// all operators
	// https://golang.org/ref/spec#Operators
	// Arithmetic operators
	// +    sum                    integers, floats, complex values, strings
	// -    difference             integers, floats, complex values
	// *    product                integers, floats, complex values
	// /    quotient               integers, floats, complex values
	// %    remainder              integers

	// Comparison operators
	// ==   equal
	// !=   not equal
	// <    less
	// <=   less or equal
	// >    greater
	// >=   greater or equal

	// Logical operators
	// &&   conditional AND
	// ||   conditional OR
	// !    NOT

	// Other operators
	// &    address of             variable
	// *    pointer to             variable
	// <-   receive                channel
	// <-   send                   channel

	// Assignment operators
	// =    simple assignment
	// +=   sum assignment
	// -=   difference assignment
	// *=   product assignment
	// /=   quotient assignment
	// %=   remainder assignment
	// &=   bitwise AND assignment
	// |=   bitwise OR assignment
	// ^=   bitwise XOR assignment
	// &^=  bit clear assignment
	// <<=  left shift assignment
	// >>=  right shift assignment

	// show a example of each operator
	var numero1 int = 10
	var numero2 int = 5

	fmt.Println("Suma: ", numero1+numero2)
	fmt.Println("Resta: ", numero1-numero2)
	fmt.Println("Multiplicacion: ", numero1*numero2)
	fmt.Println("Division: ", numero1/numero2)
	fmt.Println("Modulo: ", numero1%numero2)

	fmt.Println("Comparison operators")
	fmt.Println("==: ", numero1 == numero2)
	fmt.Println("!=: ", numero1 != numero2)
	fmt.Println("<: ", numero1 < numero2)
	fmt.Println("<=: ", numero1 <= numero2)
	fmt.Println(">: ", numero1 > numero2)
	fmt.Println(">=: ", numero1 >= numero2)

	fmt.Println("Logical operators")
	fmt.Println("&&: ", numero1 == numero2 && numero1 < numero2)
	fmt.Println("||: ", numero1 == numero2 || numero1 < numero2)
	fmt.Println("! : ", !(numero1 == numero2 || numero1 < numero2))

	fmt.Println("Other operators")
	fmt.Println("&: ", &numero1)

}
