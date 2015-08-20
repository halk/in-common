package engine

import (
	"github.com/halk/in-common/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRecommendations(t *testing.T) {
	r1 := model.Relationship{"Tester", "tester1", "Test", "test1", "tested"}
	r2 := model.Relationship{"Tester", "tester2", "Test", "test1", "tested"}
	r3 := model.Relationship{"Tester", "tester2", "Test", "test2", "tested"}
	r4 := model.Relationship{"Tester", "tester2", "Test", "test3", "tested"}
	ProcessAddEvent(&r1)
	ProcessAddEvent(&r2)
	ProcessAddEvent(&r3)
	ProcessAddEvent(&r4)

	rq := model.RecommendationRequest{"Tester", "tester1", "Test", "tested", 5}
	recommendation, err := GetRecommendations(&rq)
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
	} else {
		assert.Len(t, recommendation.Results, 2, "Unexpected size of results")
	}

	ProcessRemoveEvent(&r4)
	recommendation, err = GetRecommendations(&rq)
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
	} else {
		assert.Len(t, recommendation.Results, 1, "Unexpected size of results")
	}

	ProcessRemoveEvent(&r1)
	ProcessRemoveEvent(&r2)
	ProcessRemoveEvent(&r3)
}
