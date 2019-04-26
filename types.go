package gosf

import io "github.com/graarh/golang-socketio"

// Message - Standard message type for Socket communications
type Message struct {
	Status string `json:"status"`

	ID int `json:"id,omitempty"`

	Key     string `json:"key,omitempty"`
	Version string `json:"version,omitempty"`

	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

// Session is a data store for a single client connection
type Session struct {
	ID string
}

// Plugin is the framework interface defining a plugin
type Plugin interface {
	Activate(app *map[string]interface{})
	Deactivate(app *map[string]interface{})
	PreReceive(clientMessage *Message)
	PostReceive(clientMessage *Message)
	PreRespond(clientMessage *Message, serverMessage *Message)
	PostRespond(clientMessage *Message, serverMessage *Message)
}

// Request represents a single request over an active connection
type Request struct {
	Channel  *io.Channel
	Endpoint string
}
