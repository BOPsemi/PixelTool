package controllers

import "testing"
import "github.com/stretchr/testify/assert"

func Test_NewStartUpController(t *testing.T) {
	obj := NewStartUpController()

	assert.NotNil(t, obj)
}
