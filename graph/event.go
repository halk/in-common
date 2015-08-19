package graph

import (
	"errors"
	"github.com/halk/in-common/model"
	"github.com/jmcvetta/neoism"
)

// CreateNodesAndRelationship creates a relationship and their nodes if they do
// not exist already
func CreateNodesAndRelationship(r *model.Relationship) error {
	// connect to graph
	db := GetGraph()

	// get nodes, create if they do not exist
	subject, _, err := GetOrCreateNode(db, r.Subject, "id", neoism.Props{"id": r.SubjectID})
	if err != nil {
		panic(err)
	}
	if subject == nil || subject.Id() == 0 {
		return errors.New("Could not create subject node")
	}

	object, _, err := GetOrCreateNode(db, r.Object, "id", neoism.Props{"id": r.ObjectID})
	if err != nil {
		panic(err)
	}
	if object == nil || object.Id() == 0 {
		return errors.New("Could not create object node")
	}

	return CreateOrIncRelationship(db, r)
}

// RemoveRelationship removes a relationship
func RemoveRelationship(r *model.Relationship) error {
	return DeleteRelationship(GetGraph(), r)
}
