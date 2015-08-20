package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpperCaseFirst(t *testing.T) {
	assert.Equal(t, "Test", UpperCaseFirst("test"), "Unexpected result")
	assert.Equal(t, "Test", UpperCaseFirst("Test"), "Unexpected result")
	assert.Equal(t, "", UpperCaseFirst(""), "Unexpected result")
}
