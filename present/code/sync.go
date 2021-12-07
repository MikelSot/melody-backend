package main

import (
	"fmt"
	"sync"
)

//START OMIT
func main() {
	var wg sync.WaitGroup // HL

	sayHello := func() {
		defer wg.Done() // HL
		fmt.Println("hello")
	}
	wg.Add(1) // HL

	go sayHello()
	wg.Wait() // HL
}

//END OMIT
