package main

import "fmt"

// pointer help to save memory instead of copying you make refrence to the memory address
//Pointers are usually preferred while passing a struct as an argument or while declaring a method for a defined type.
// & - to refer to the memory address
// * - go get the value form the meomery address
func incremant(i *int) {
	*i++
}

func main() {
	i := 10
	incremant(&i)
	fmt.Println(i)
}
