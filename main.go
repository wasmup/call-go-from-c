package main

import "C"
import "fmt"

// Add returns a+b
//export Add
func Add(a, b uint64) uint64 {
	return a + b
}

// Hello returns CString
//export Hello
func Hello() *C.char {
	return C.CString("Hello from Go func: Hello().")
}

func main() {
	fmt.Println("Do not run this. run 'make' instead.") // fun
}
