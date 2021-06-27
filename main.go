package main

import "fmt"

const korean = "Korean"
const french = "French"

const englishHelloPrefix = "Hello, "
const englishDefaultName = "World"

const koreanHelloPrefix = "안녕하세요, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = englishDefaultName
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case korean:
		prefix = koreanHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main()  {
	fmt.Println(Hello("Seungmin", ""))
}