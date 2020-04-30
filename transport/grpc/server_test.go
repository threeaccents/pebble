package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/threeaccents/cache/mock"
)

const testPort = ":5555"

func TestNewServer(t *testing.T) {
	s := NewServer(&mock.CacheStorage{})

	assert.NotNil(t, s.Storage, "storage should be set")
	assert.Equal(t, defaultPort, s.Port, "default port be set")
}

func TestNewServerWithOptions(t *testing.T) {
	s := NewServer(
		&mock.CacheStorage{},
		ServerPort(testPort),
	)

	assert.NotNil(t, s.Storage, "storage should be set")
	assert.Equal(t, testPort, s.Port, "test port be set")
}
