package kondo

import (
	"errors"
	"fmt"
	"github.com/kelindar/binary"
	"github.com/olahol/melody"
	"reflect"
)

var (
	typeRegistry = make(map[string]reflect.Type)

	// Connect and disconnect callback arrays are responsible for handling connect and disconnect events.
	// These are separate, because they do not take a dynamic type.
	connectCallbacks    []func(sender *melody.Session)
	disconnectCallbacks []func(sender *melody.Session)

	callbacks = make(map[reflect.Type][]any)
)

// OnConnect adds a callback to call whenever a session connects to the server.
// Note: sender will be nil in client callbacks.
func OnConnect(callback func(sender *melody.Session)) {
	connectCallbacks = append(connectCallbacks, callback)
}

// OnDisconnect adds a callback to call whenever a session disconnects from the server.
// Note: sender will be nil in client callbacks.
func OnDisconnect(callback func(sender *melody.Session)) {
	disconnectCallbacks = append(disconnectCallbacks, callback)
}

// On adds a callback to be called whenever the specified message type T is received.
// Note: sender will be nil in client callbacks.
func On[T any](callback func(sender *melody.Session, message T)) {
	handlerType := reflect.TypeOf(callback).In(1)

	// Register the type in the type registry.
	typeRegistry[handlerType.String()] = handlerType

	// Add the callback to the router.
	// So we can reference it when processing messages.
	callbacks[handlerType] = append(callbacks[handlerType], callback)
}

func processMessage(sender *melody.Session, nm netMessage) error {
	t := typeRegistry[nm.Ref]

	if t == nil {
		return errors.New(fmt.Sprintf("message of type %s was not registered", nm.Ref))
	}

	// Creates an instance of the T (message value) type.
	instance := reflect.New(t)

	// Unmarshal the NetMessage Data field as the message value type.
	err := binary.Unmarshal(nm.Data, instance.Interface())
	if err != nil {
		return err
	}

	// Call all the registered callbacks.
	amtCallbacks := len(callbacks[t])
	for _, callback := range callbacks[t] {
		callbackValue := reflect.ValueOf(callback)

		// Build our arguments for the callback itself.
		// Fills the first argument with the sender, and the second with the fina l unmarshaled struct.
		arguments := []reflect.Value{reflect.ValueOf(sender), instance.Elem()}

		// Calls the callback, delivering the data back to the .On function.
		callbackValue.Call(arguments)
	}

	if amtCallbacks == 0 {
		return errors.New("no callbacks were registered for type " + nm.Ref)
	}

	return nil
}

func callConnect(sender *melody.Session) {
	for _, callback := range connectCallbacks {
		callback(sender)
	}
}

func callDisconnect(sender *melody.Session) {
	for _, callback := range disconnectCallbacks {
		callback(sender)
	}
}
