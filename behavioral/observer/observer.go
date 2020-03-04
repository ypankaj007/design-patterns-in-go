package main

// Observer is an interface
// Who needs to nofify will implement
type Observer interface {
	// OnNotify and event will observe
	OnNotify(Event)
}
