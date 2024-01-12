package main

import "fmt"

func main() {
	//all data types
	var nombre string = "Juan"
	var apellido string = "Perez"
	var edad int = 25
	var altura float32 = 1.75
	var casado bool = false
	var birthdate = "1990-01-01"
	var birthdate2 = "1990-01-01 12:00:00"
	var birthdate3 = "1990-01-01 12:00:00 -0500"
	var birthdate4 = "1990-01-01 12:00:00 -0500 -05"
	var birthdate5 = "1990-01-01 12:00:00 -0500 -05 CST"
	var birthdate6 = "1990-01-01 12:00:00 -0500 -05 CST America/Bogota"

	//other data types
	var numero1 uint8 = 255
	var numero2 uint16 = 65535
	var numero3 uint32 = 4294967295
	var numero4 uint64 = 18446744073709551615
	var numero5 int8 = -128
	var numero6 int16 = -32768
	var numero7 int32 = -2147483648
	var numero8 int64 = -9223372036854775808
	var numero9 float64 = 9223372036854775808
	var numero10 complex64 = 5 + 5i
	var numero11 complex128 = 5 + 5i
	var numero12 byte = 255
	var numero13 rune = 2147483647
	var numero14 rune = -2147483648

	//give me a data types list with description
	//https://golang.org/ref/spec#Numeric_types
	// 	uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	//print all
	println(nombre, apellido, edad, altura, casado, birthdate, birthdate2, birthdate3, birthdate4, birthdate5, birthdate6)

	//print all
	println(numero1, numero2, numero3, numero4, numero5, numero6, numero7, numero8, numero9, numero10, numero11, numero12, numero13, numero14)

	//  print numero4 with comma and dot 000,000.00
	//  print numero4 with comma and dot 000.000,00
	//  print numero4 with comma and dot 000.000.000,00
	//  print numero4 with comma and dot

	// Print numero4 with comma and dot formatting
	formatted := fmt.Sprintf("%'d", numero4)
	formatted = formatted[:3] + "," + formatted[3:6] + "." + formatted[6:]
	fmt.Println(formatted)

}
