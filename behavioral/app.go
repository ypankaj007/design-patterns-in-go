package main

import (
	"bufio"
	ob "design-patterns-in-go/behavioral/observer"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Observer application .......")
	// Initialize notifier
	privatePortal := ob.Initialize()
	fmt.Println("Registering linkedin observer")
	privatePortal.Attach(&ob.Linkedin{})

	fmt.Println("Registering facebook observer")
	privatePortal.Attach(&ob.Facebook{})

	fmt.Println("Registering twitter observer")

	privatePortal.Attach(&ob.Twitter{})

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		e := ob.Event{Message: scanner.Text()}
		privatePortal.Notify(e)
	}
	fmt.Println(scanner.Err())
}
