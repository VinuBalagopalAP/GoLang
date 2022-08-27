package main

import "fmt"

func type_conversion() {

	/// [ Declare a type ]
	type fahrenhiet int
	type celsius int

	var f fahrenhiet = 32
	var c celsius = 0

	fmt.Println(f, c)

	/// [ Convert a type ]
	/// [ Equivalent to fahrenhiet(32) ]
	c = celsius((f - 32) * 5 / 9)

	fmt.Println(c)

}
