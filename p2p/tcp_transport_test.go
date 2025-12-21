package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAdderess := ":4000"
	transport := NewTCPTransport(listenAdderess)
	assert.Equal(t, transport.listenAdderess, listenAdderess)

	// Server start/stop tests can be added here in the future
	assert.Nil(t, transport.ListenAndAccept())
}
