package winkmqtt

// Device model and the Device store (also need the device communicator)

// Device represents a given device in the wink system
type Device struct {
	ID int
	// Array of attributes?
}

// Do we need to have a reader and writer?
// Can apron do both?
// Can we have a reader and a writer compiled into a service?

// DeviceService is an implementable interface
type DeviceService interface {
}
