package hello

const (
	spanish    = "Spanish"
	french     = "French"
	portuguese = "Portuguese"

	englishPrefix    = "Hello, "
	spanishPrefix    = "Hola, "
	frenchPrefix     = "Bonjour, "
	portuguesePrefix = "Olá, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case portuguese:
		prefix = portuguesePrefix
	default:
		prefix = englishPrefix
	}

	return
}
