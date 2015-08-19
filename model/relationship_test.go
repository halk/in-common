package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebug(t *testing.T) {
	r := Relationship{"Tester", "tester1", "Test", "test1", "tested"}
	assert.Equal(t, "Tester(tester1)-TESTED->Test(test1)", r.Debug(), "Unexpected result")
}
