package iter

const repeatCount = 5

func Repeat(a string) string {
	var repeates string
	for i := 0; i < repeatCount; i++ {
		repeates += a
	}
	return repeates
}
