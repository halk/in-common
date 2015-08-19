package engine

import (
	"inCommon/graph"
	"inCommon/model"
)

// Return recommendations
func GetRecommendations(rq *model.RecommendationRequest) (*model.Recommendation, error) {
	return graph.GetRecommendations(rq)
}
