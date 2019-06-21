
# Observer Design Pattern in GO

The observer pattern is a software design pattern in which an object, called the subject, maintains a list of its dependents, called observers, and notifies them automatically of any state changes, usually by calling one of their methods.

# Example

 - In the code, I took a simple message posting example over social media (Linkedin, Facebook and Twitter).
 -  User will input the message by command line and it will observer in all social media portal.
 
# Implementation

  - import observer package in our application 
  - Create Event and Register Observer
  - Resiter observer 
    ```sh
    privatePortal.Attach(observer) // observer defines Observer object 

    ```

- Remove observer 
    ```sh
    privatePortal.Detach(observer) // observer defines Observer object

    ```

    - Notify to observer 
    ```sh
    privatePortal.Notify(e) // e defines Event object

    ```
  
```sh
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

```