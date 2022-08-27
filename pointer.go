package main

import "fmt"

func pointer() {

	/// [ Declaring x as int in short declaration ]
	x := 1
	/// [ Declaring p as the pointer of x in short declaration  ]
	p := &x

	/// [ Print the address of x ]
	fmt.Println(p)

	/// [ Print the value of x ]
	fmt.Println(*p)
}
