package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoutes(t *testing.T) {
	assert.IsType(t, Routes{}, routes, "Unexpected type of routes")
	assert.Len(t, routes, 4, "Unexpected number of routes")
	assert.Equal(t, "Index", routes[0].Name, "Unexpected route name")
	assert.Equal(t, "AddEvent", routes[1].Name, "Unexpected route name")
	assert.Equal(t, "RemoveEvent", routes[2].Name, "Unexpected route name")
	assert.Equal(t, "Recommend", routes[3].Name, "Unexpected route name")
}
