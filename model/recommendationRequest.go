package model

// RecommendationRequest holds all details required to get a recommendation
type RecommendationRequest struct {
	Subject      string `schema:"subject"`
	SubjectID    string `schema:"subject_id"`
	Object       string `schema:"object"`
	Relationship string `schema:"relationship"`
	Limit 		   int `schema:"limit"`
}
