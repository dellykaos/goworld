package goworld_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert.Equal(t, "Hello, world!", goworld.Hello(""))
	assert.Equal(t, "Hello, Budi!", goworld.Hello("Budi"))
}
