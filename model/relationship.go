package model

import (
	"fmt"
	"strings"
)

// Relationship represents a Neo4j relationship
type Relationship struct {
	Subject      string `json:"subject"`
	SubjectID    string `json:"subject_id"`
	Object       string `json:"object"`
	ObjectID     string `json:"object_id"`
	Relationship string `json:"relationship"`
}

// Debug returns a debug presentation of a relationship
func (r *Relationship) Debug() string {
	return fmt.Sprintf(
		"%s(%s)-%s->%s(%s)", r.Subject, r.SubjectID,
		strings.ToUpper(r.Relationship), r.Object, r.ObjectID,
	)
}
