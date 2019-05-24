package main

import "fmt"

func main() {
	// array in go have a fixed lenght
	var arr [3]int
	arr[0] = 0
	arr[1] = 3
	arr[2] = 4
	fmt.Println(arr)
	// declaring any arry
	list := [4]string{"Sigo", "Nibraz", "Musa", "Benito"}
	fmt.Println(list)
	// two dimensional array
	set := [3][3]int{{2, 3, 4}, {4, 5, 7}, {8, 9, 6}}
	fmt.Println(set)
	// slices
	//- almost like an are but you don't set length because is extendable
	slice := []int{2, 3, 4, 6}
	newSLice := append(slice, 7) // add item to the array
	// when you add item to an array go create a new array for you
	fmt.Println(slice)
	fmt.Println(newSLice)
	// slicing a slice
	fmt.Println(newSLice[:3])
	// map just like js object which have a key and value
	// init
	n := map[string]string{"name": "sigo"}
	fmt.Println(n)
	test := make(map[string]int)
	test["english"] = 78
	test["math"] = 80
	test["PE"] = 88
	fmt.Println(test)
	delete(test, "PE")
	fmt.Println(test)
}
