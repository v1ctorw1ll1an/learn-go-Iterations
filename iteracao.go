package iteracao

// Repeat returns a string of s repeated n times
func Repeat(s string, n int) string {
	temp := ""
	for range n {
		temp += s
	}

	return temp
}
