
package main

import "fmt"

// function name is: squares
// function return value is: func() int
func squares() func() int{
    var x int
    return func() int {
        x++
        return x * x
    }
}

func main() {
	f := squares()  // after squares returns, x is hidden inside f
	fmt.Printf("type of f is: %T\n", f)  // func() int

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// try to call a function literal
	func() {
		fmt.Println("Call a function literal")
	}()

	fmt.Println(f())
}