package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

func startEmulator() {
	// Create a new RAM object
	ram := RAM{}.initMemory()
	fmt.Println(ram)
}
