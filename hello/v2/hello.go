package v2

func Hello(name string) string {

	if name == "" {
		return "Hello, World"
	}

	return "Hello, " + name
}
