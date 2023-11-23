package iteration

func Repeat(character string, n int) string {
	var repeat string

	for i := 0; i < n; i++ {
		repeat += character
	}
	return repeat
}
