package event

// Listener ..
// All custom event listeners must satisfy this Listener interface.
type Listener interface {
	Listen(event interface{})
}
