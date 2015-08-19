// The inCommon/model package has data structs and minimal data manipulating methods
package model

// Recommendation contains a list of recommendations
type Recommendation struct {
	Results map[string]float32 `json:"results"`
}

func NewRecommendation() *Recommendation {
	r := new(Recommendation)
	r.Results = map[string]float32{}
	return r
}

// Add adds a recommendation to the result list
func (r *Recommendation) Add(ID string, weight float32) {
	r.Results[ID] = weight
}
