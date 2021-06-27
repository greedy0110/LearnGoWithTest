package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishDefaultName = "World"

func Hello(name string) string {
	if name == "" {
		name = englishDefaultName
	}
	return englishHelloPrefix + name
}

func main()  {
	fmt.Println(Hello("Seungmin"))
}