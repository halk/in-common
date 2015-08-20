package engine

import (
	"github.com/halk/in-common/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessAddEvent(t *testing.T) {
	r := model.Relationship{"Tester", "tester1", "Test", "test1", "tested"}
	if err := ProcessAddEvent(&r); err != nil {
		assert.Fail(t, "Unexpected error:"+err.Error())
	}
}

func TestProcessRemoveEvent(t *testing.T) {
	r := model.Relationship{"Tester", "tester1", "Test", "test1", "tested"}
	if err := ProcessRemoveEvent(&r); err != nil {
		assert.Fail(t, "Unexpected error:"+err.Error())
	}
}
