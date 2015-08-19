package api

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter()
	assert.IsType(t, mux.NewRouter(), router, "Unexpected type of router")
}
