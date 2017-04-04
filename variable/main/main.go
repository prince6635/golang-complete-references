package main

import (
	"github.com/golang-complete-references/variable/variableutil"
	"fmt"
)

func main() {
	// declare, assign, initialize
	var msg string
	msg = "Hello, world!"
	fmt.Println(msg)

	variableutil.Shorthand()
	variableutil.VarZeroValue()
	variableutil.All()
}
