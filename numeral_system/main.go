package main

import "fmt"

func main() {
	num := 42
	// %d: decimal; %b: binary; %x: hex
	fmt.Printf("%d - %b - %x \n", num, num, num)
	fmt.Printf("%d - %b - %#x \n", num, num, num)
	fmt.Printf("%d - %b - %#X \n", num, num, num)
	fmt.Printf("%d \t %b \t %#X \n", num, num, num)

	for i := 0; i < 100; i++ {
		fmt.Printf("%d \t %b \t %#X \n", i, i, i)

		// http://stackoverflow.com/questions/700187/unicode-utf-ascii-ansi-format-differences
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
	}
}