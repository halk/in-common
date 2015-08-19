package engine

import (
	"github.com/halk/in-common/graph"
	"github.com/halk/in-common/model"
)

// Return recommendations
func GetRecommendations(rq *model.RecommendationRequest) (*model.Recommendation, error) {
	return graph.GetRecommendations(rq)
}
