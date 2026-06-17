package main

import (
	"fmt"
	"strings"
)

const Spanish = "SPANISH"
const French = "FRENCH"
const Esperanto = "ESPERANTO"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const esperantoHelloPrefix = "Saluton, "

func Hello(name string, language string) string {

	var prefix string

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

	if name == "" {
		return prefix + "World"
	}

	return prefix + name
}
func main() {
	fmt.Println(Hello("World", ""))
}
