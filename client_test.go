package kondo

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(url.URL{Scheme: "ws", Host: "localhost:8991", Path: "/ws"})
	assert.NotNil(t, c)
	assert.NotNil(t, c.m)
}
