package main

func funtionWithParameters(num1, num2 int) int {
	return num1 + num2

}

func functionWithReturn() int {
	return 1
}

func functionWithoutReturn() {
	println("Hello world")

}

func functionWithMultipleReturn() (int, int) {
	other := func(v int) int {
		return v
	}
	return 1, other(2)
}

func functionWithNamedReturn() (x int) {
	x = 1
	return
}

func main() {

	functionWithoutReturn()
	println(functionWithReturn())
	println(funtionWithParameters(1, 2))
	println(functionWithMultipleReturn())
	println(functionWithNamedReturn())

	//how to use a functionWithMultipleReturn
	var x, y int = functionWithMultipleReturn()
	println(x, y)

}
