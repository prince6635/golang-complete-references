package stringutil

// this demonstrates how an unexported function
// can be used by an exported function in the same package
func reverseInternal(s string) string {
	chars := []rune(s)

	for i, j := 0, len(chars)-1; i < len(chars)/2; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}
