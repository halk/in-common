package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRecommendationRequest(t *testing.T) {
	rq := RecommendationRequest{"Tester", "test1", "Test", "tested", 2}
	assert.IsType(t, RecommendationRequest{}, rq, "Unexpected type for recommendation request struct")
	assert.Equal(t, "Tester", rq.Subject, "Unexpected value for subject")
	assert.Equal(t, "test1", rq.SubjectID, "Unexpected value for subject ID")
	assert.Equal(t, "Test", rq.Object, "Unexpected value for object")
	assert.Equal(t, "tested", rq.Relationship, "Unexpected value for relationship")
	assert.Equal(t, 2, rq.Limit, "Unexpected value for limit")
}
