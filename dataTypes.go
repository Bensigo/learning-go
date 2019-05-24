package main

import "fmt"

func main() {
	// data type in go
	// go have numeric , string, boolean, float

	/**
		------------------------------
		numeric type | range
		------------------------------
		| int8			 | -128 -127			|
		| uint 8  	 | 0 to 127       |
		_______________________________
		| int 16		 | -2^15 - 2^15-1 |
		| uint 16  	 | 0 to 2^15-1    |
		_______________________________
		| int 32		 | -2^31 - 2^32-1 |
	  | uint 32  	 | 0 to 2^32-1    |
		_______________________________
		| int 64		 | -2^63 - 2^63-1 |
		| uint 64  	 | 0 to 2^63-1    |
		_______________________________
		| int    		 | platform 			|
		|						 | dependent 			|
		_______________________________
	**/

	// numeric data type
	var small int8 // range -128 to 127
	small = 8
	fmt.Println(small)
	// int a platform depentent numeric type
	var num int = 9
	fmt.Println(num)
	var myInt16 int16 = 1000 // range -2**15 to 2**15 -1
	fmt.Println(myInt16)

	var myInt32 int32 = 500000 // range -2**31 to 2**31 -1
	fmt.Println(myInt32)
	// int64 range	-2**63 to 2**63

	// string are double qoutes
	var name = "my name is Sigo"

	// multiline string
	const note string = `
		this is a multiline string
	`
	// boolean
	var isValid bool = false

	// floating number main use for money value which we have  float32 and float64
	const amount float32 = 100.89
	const total float64 = 5000.83738
}
