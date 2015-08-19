package graph

import (
	"github.com/stretchr/testify/assert"
	"inCommon/model"
	"testing"
)

func TestCreateNodesAndRelationship(t *testing.T) {
	r := model.Relationship{"graphtest3_1", "node1", "graphtest3_2", "node2", "tested"}
	if err := CreateNodesAndRelationship(&r); err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
	}
}

func TestRemoveRelationship(t *testing.T) {
	r := model.Relationship{"graphtest3_1", "node1", "graphtest3_2", "node2", "tested"}
	if err := CreateNodesAndRelationship(&r); err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
	}
	if err := RemoveRelationship(&r); err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
	}
}
