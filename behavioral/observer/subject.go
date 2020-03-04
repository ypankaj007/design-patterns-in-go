package main

// Subject defines notifier
// Only implemented by who wish to nofity
type Subject interface {
	// Attach Observer
	// An Observer will register
	Attach(Observer)

	// Detach an registered observer
	Detach(Observer)

	// Notify will send change event
	Notify(Event)
}
