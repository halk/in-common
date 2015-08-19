// The in-common/engine package provides the internal interface to the graph package
package engine

import (
	"github.com/halk/in-common/graph"
	"github.com/halk/in-common/model"
)

// Process add events
func ProcessAddEvent(r *model.Relationship) error {
	return graph.CreateNodesAndRelationship(r)
}

// Process remove events
func ProcessRemoveEvent(r *model.Relationship) error {
	return graph.RemoveRelationship(r)
}
