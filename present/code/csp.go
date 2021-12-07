package main

import "fmt"

//START OMIT
func main() {
	sayHello := make(chan string, 2) // HL

	go func() {
		sayHello <- "Hello" // HL
	}()

	fmt.Println(<-sayHello) // HL
}

//END OMIT
