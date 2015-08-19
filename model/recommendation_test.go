package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRecommendation(t *testing.T) {
	r := NewRecommendation()
	assert.IsType(t, Recommendation{}, *r, "Unexpected type for recommendation struct")
}

func TestAdd(t *testing.T) {
	r := NewRecommendation()
	assert.Len(t, r.Results, 0, "Unexpected size of Results")
	r.Add("test1", 1)
	assert.Len(t, r.Results, 1, "Unexpected size of Results")
	assert.Equal(t, float32(1), r.Results["test1"], "Unexpected value for element")
	r.Add("test2", 0.3332)
	assert.Len(t, r.Results, 2, "Unexpected size of Results")
	assert.Equal(t, float32(0.3332), r.Results["test2"], "Unexpected value for element")
}
