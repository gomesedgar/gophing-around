package helloworld

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + exclamationPoint
}

func greetingPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	case german:
		return germanHelloPrefix
	default:
		return englishHelloPrefix
	}
}
