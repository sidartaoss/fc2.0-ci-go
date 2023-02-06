package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	// arrange
	a := 1
	b := 2
	expected := 3
	// act
	actual := Soma(a, b)

	// assert
	assert.Equal(t, expected, actual)
	// fc2.0-ci-go
}
