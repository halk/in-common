package engine

import (
	"github.com/stretchr/testify/assert"
	"inCommon/model"
	"testing"
)

func TestProcessAddEvent(t *testing.T) {
	r := model.Relationship{"Tester", "tester1", "Test", "test1", "tested"}
	if err := ProcessAddEvent(&r); err != nil {
		assert.Fail(t, "Unexpected error")
	}
}

func TestProcessRemoveEvent(t *testing.T) {
	r := model.Relationship{"Tester", "tester1", "Test", "test1", "tested"}
	if err := ProcessRemoveEvent(&r); err != nil {
		assert.Fail(t, "Unexpected error")
	}
}
