// The inCommon/engine package provides the internal interface to the graph package
package engine

import (
	"inCommon/graph"
	"inCommon/model"
)

// Process add events
func ProcessAddEvent(r *model.Relationship) error {
	return graph.CreateNodesAndRelationship(r)
}

// Process remove events
func ProcessRemoveEvent(r *model.Relationship) error {
	return graph.RemoveRelationship(r)
}
