package winkmqtt

// Message is a struct that represents a message passed
// in from a given message broker
type Message struct {
	Topic   string
	Payload string
}

// MessageService is an implementable interface for any given message broker
type MessageService interface {
}
