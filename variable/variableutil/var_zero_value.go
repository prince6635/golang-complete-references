package variableutil

import "fmt"

func VarZeroValue()  {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("(%T) %v \n", a, a)
	fmt.Printf("(%T) %v \n", b, b)
	fmt.Printf("(%T) %v \n", c, c)
	fmt.Printf("(%T) %v \n", d, d)
}
