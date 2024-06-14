package main

import "fmt"

const (
	portuguese = "portuguese"
	french     = "french"

	portuguesePrefix = "Ola, "
	frenchPrefix     = "Bonjour, "
	englishPrefix    = "Hello, "
)

func getGreetingPrefixes() (greetings map[string]string) {
	greetings = map[string]string{
		portuguese: portuguesePrefix,
		french:     frenchPrefix,
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	greetingPrefixes := getGreetingPrefixes()
	greeting, greetingExists := greetingPrefixes[language]

	if greetingExists {
		prefix = greeting
		return
	}
	prefix = englishPrefix

	return
}

func main() {
	fmt.Print(Hello("world", ""))
}
