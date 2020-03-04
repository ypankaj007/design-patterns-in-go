package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Observer application .......")
	// Initialize notifier
	privatePortal := Initialize()
	fmt.Println("Registering linkedin observer")
	privatePortal.Attach(&Linkedin{})

	fmt.Println("Registering facebook observer")
	privatePortal.Attach(&Facebook{})

	fmt.Println("Registering twitter observer")

	privatePortal.Attach(&Twitter{})

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		e := Event{Message: scanner.Text()}
		privatePortal.Notify(e)
	}
	fmt.Println(scanner.Err())
}
