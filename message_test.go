package kondo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ChatMessageExample struct {
	Message string
	Sender  string
	Private bool
}

var nm = newNetMessage(ChatMessageExample{
	Message: "This is a test",
	Sender:  "Tester Test",
	Private: false,
})

func TestEncode(t *testing.T) {
	encoded, err := encode(nm)

	assert.Nil(t, err)
	assert.NotNil(t, encoded)
}

func TestParseIncoming(t *testing.T) {
	encoded, err := encode(nm)

	assert.Nil(t, err)
	assert.NotNil(t, encoded)

	nm, err := parseIncoming(encoded)
	assert.Nil(t, err)
	assert.NotNil(t, nm)
}
