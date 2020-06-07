package aenetwork

// Represents an object that implements network listening capabilities.
type NetworkListener interface {
	// Called when a client connects.
	OnConnect(connection *ClientConnection)

	// Called when a client disconnects.
	OnDisconnect(connection *ClientConnection)
}
