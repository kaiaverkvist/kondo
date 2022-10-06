package kondo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer(":8991", "/ws")
	assert.NotNil(t, s)
	assert.NotNil(t, s.m)
}
