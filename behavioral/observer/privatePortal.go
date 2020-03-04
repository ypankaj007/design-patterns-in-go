package main

// PrivatePortal defines a user's portal
// where user will add remove the SocialMedia portals before sending posting message
type PrivatePortal struct {
	// Using a map with an empty struct allows to keep the observers
	observers map[Observer]struct{}
}

// Attach: registering observer
func (o *PrivatePortal) Attach(observer Observer) {
	o.observers[observer] = struct{}{}
}

// Implementing Detach
func (o *PrivatePortal) Detach(observer Observer) {
	delete(o.observers, observer)
}

// Implementing Notify
func (p *PrivatePortal) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}

// Initialize observer
func Initialize() *PrivatePortal {
	return &PrivatePortal{
		observers: map[Observer]struct{}{},
	}
}
