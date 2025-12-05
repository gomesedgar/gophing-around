package main

func main() {
	println(HelloWorld())
}

const englishHelloPrefix = "Hello, "
const exclamationPoint = "!"

func HelloWorld() string {
	return Hello("world")
}

func Hello(name string) string {
	return englishHelloPrefix + name + exclamationPoint
}
