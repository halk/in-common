package graph

import (
	"github.com/halk/in-common/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRecommendations_Empty(t *testing.T) {
	rq := model.RecommendationRequest{"Tester", "graphtest4_1", "graphtest4", "tested", 10}
	recommendation, err := GetRecommendations(&rq)
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
		return
	}

	assert.IsType(t, model.Recommendation{}, *recommendation, "Unexpected type")
	assert.Len(t, recommendation.Results, 0, "Unexpected length")
}

func TestGetRecommendations(t *testing.T) {
	r1 := model.Relationship{"Tester", "graphtest4_1", "graphtest4", "node1", "tested"}
	CreateNodesAndRelationship(&r1)
	r2 := model.Relationship{"Tester", "graphtest4_2", "graphtest4", "node1", "tested"}
	CreateNodesAndRelationship(&r2)
	r3 := model.Relationship{"Tester", "graphtest4_2", "graphtest4", "node2", "tested"}
	CreateNodesAndRelationship(&r3)

	rq := model.RecommendationRequest{"Tester", "graphtest4_1", "graphtest4", "tested", 10}
	recommendation, err := GetRecommendations(&rq)
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
		return
	}

	assert.IsType(t, model.Recommendation{}, *recommendation, "Unexpected type")
	assert.Len(t, recommendation.Results, 1, "Unexpected length")
	assert.Equal(t, float32(1), recommendation.Results["node2"])

	r4 := model.Relationship{"Tester", "graphtest4_2", "graphtest4", "node3", "tested"}
	CreateNodesAndRelationship(&r4)
	r5 := model.Relationship{"Tester", "graphtest4_3", "graphtest4", "node1", "tested"}
	CreateNodesAndRelationship(&r5)
	r6 := model.Relationship{"Tester", "graphtest4_3", "graphtest4", "node2", "tested"}
	CreateNodesAndRelationship(&r6)

	recommendation, err = GetRecommendations(&rq)
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
		return
	}

	assert.Len(t, recommendation.Results, 2, "Unexpected length")
	assert.Equal(t, float32(2), recommendation.Results["node2"])
	assert.Equal(t, float32(1), recommendation.Results["node3"])

	rq.Limit = 1
	recommendation, err = GetRecommendations(&rq)
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
		return
	}

	assert.Len(t, recommendation.Results, 1, "Unexpected length")
	assert.Equal(t, float32(2), recommendation.Results["node2"])

	RemoveRelationship(&r3)
	RemoveRelationship(&r6)
	recommendation, err = GetRecommendations(&rq)
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
		return
	}

	assert.Len(t, recommendation.Results, 1, "Unexpected length")
	assert.Equal(t, float32(1), recommendation.Results["node3"])

	RemoveRelationship(&r1)
	RemoveRelationship(&r2)
	RemoveRelationship(&r4)
	RemoveRelationship(&r5)
}
