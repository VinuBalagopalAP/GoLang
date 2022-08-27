package main

import "fmt"

func main() {
	/// [Print Hello World]
	fmt.Println("Hello, World!")

	/// [declare and initialize a variable]
	/// [ Syntax: var <variable_name> <data_type> = <value> ]
	var integer int = 1

	fmt.Println(integer)

	/// [ declare and initialize a variable with multiple values ]
	/// [ Syntax: var <variable_name> <data_type> = <value1>, <value2>, <value3> ]
	var integer1, integer2, integer3 int = 1, 2, 3

	fmt.Println(integer1, integer2, integer3)

	/// [ Can also declare and initialize a variable and string at the same time ]
	/// [ Syntax: var <variable_name> <data_type> = <value> <string> ]
	var integer4, string1 = 4, "Hello"

	fmt.Println(integer4, string1)

	/// [ short declaration of a boolean variable ]
	/// [ Syntax: var <variable_name> <data_type> = <value> ]
	var boolean bool = true

	fmt.Println(boolean)
}

///  To run the above code, type the following in the terminal:[ go run main.go ]