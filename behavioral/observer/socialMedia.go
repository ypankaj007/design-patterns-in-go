package observer

import "fmt"

// Linkedin social media portal
type Linkedin struct {
	// Message on linkedin
	Message string
}

// Implementing OnNotify by linkedin to receive the event
func (lin *Linkedin) OnNotify(e Event) {

	// Meesage on facebook
	lin.Message = e.Message
	fmt.Println("Linkedin received message:  ", lin.Message)
}

// Facebook portal
type Facebook struct {
	Message string
}

// Implementing OnNotify by facebook to receive the event
func (fb *Facebook) OnNotify(e Event) {
	fb.Message = e.Message
	fmt.Println("Facebook received message:  ", fb.Message)
}

// Twitter
type Twitter struct {
	Message string
}

// mplementing OnNotify by twitter to receive the event
func (t *Twitter) OnNotify(e Event) {
	t.Message = e.Message
	fmt.Println("Twitter received message:  ", t.Message)
}
