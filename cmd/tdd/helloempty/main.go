package main

import (
	"fmt"
	"strings"
)

const (
	Spanish              = "SPANISH"
	French               = "FRENCH"
	Esperanto            = "ESPERANTO"
	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchHelloPrefix    = "Bonjour, "
	esperantoHelloPrefix = "Saluton, "
)

func Hello(name string, language string) (prefix string) {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch strings.ToUpper(language) {
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	case Esperanto:
		prefix = esperantoHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
