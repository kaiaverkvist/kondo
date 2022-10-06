package kondo

import (
	"github.com/olahol/melody"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ExampleChatMessage struct {
	message string
}

func Test_RouterOn(t *testing.T) {
	encoded, err := encode(newNetMessage(ExampleChatMessage{message: "test message!"}))

	assert.Nil(t, err)
	assert.NotNil(t, encoded)

	var called bool
	On[ExampleChatMessage](func(sender *melody.Session, message ExampleChatMessage) {
		called = true
	})

	nm, err := parseIncoming(encoded)
	assert.Nil(t, err)
	assert.NotNil(t, nm)

	err = processMessage(&melody.Session{}, nm)
	assert.Nil(t, err)
	assert.True(t, called)
}
