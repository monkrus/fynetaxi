package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	n := &note{content: "This note"}
	assert.Equal(t, "This note", n.title())

	n = &note{content: "line1\nline2"}
	assert.Equal(t, "line1", n.title())
}
