// The in-common/graph package holds the actual graph and recommendation logic
package graph

import (
	"fmt"
	"github.com/halk/in-common/model"
	"github.com/halk/in-common/util"
	"github.com/jmcvetta/neoism"
	"os"
	"strings"
)

// GetGraph returns a connected neoism.Database
func GetGraph() *neoism.Database {
	if os.Getenv("IN_COMMON_NEO4J_DSN") != "" {
		dsn := os.Getenv("IN_COMMON_NEO4J_DSN")
	} else {
		dsn := "http://neo4j:vagrant@localhost:7474/db/data"
	}

	db, err := neoism.Connect(dsn)
	if err != nil {
		panic(err)
	}

	return db
}

// GetOrCreateNode returns a Node, creates one if it does not exist. This method
// overrides the neoism one as there is a bug
func GetOrCreateNode(db *neoism.Database, label, key string, p neoism.Props) (n *neoism.Node, created bool, err error) {
	node, created, err := db.GetOrCreateNode(util.UpperCaseFirst(label), key, p)
	// due to a bug in neoism, label is not added
	// see: https://github.com/jmcvetta/neoism/issues/62
	if created {
		node.AddLabel(util.UpperCaseFirst(label))
	}
	return node, created, err
}

// CreateOrIncRelationship creates a relationship or increments a property if it
// already exists
func CreateOrIncRelationship(db *neoism.Database, r *model.Relationship) error {
	cq := neoism.CypherQuery{
		Statement: fmt.Sprintf(
			`
			MATCH (a:%s), (b:%s)
			WHERE a.id = {subjectId} AND b.id = {objectId}
			CREATE UNIQUE (a)-[r:%s]->(b)
			SET r.frequency = COALESCE(r.frequency, 0) + 1
			`,
			util.UpperCaseFirst(r.Subject), util.UpperCaseFirst(r.Object),
			strings.ToUpper(r.Relationship),
		),
		Parameters: neoism.Props{
			"subjectId": r.SubjectID, "objectId": r.ObjectID,
		},
	}
	return db.Cypher(&cq)
}

// DeleteRelationship deletes a relationship
func DeleteRelationship(db *neoism.Database, r *model.Relationship) error {
	cq := neoism.CypherQuery{
		Statement: fmt.Sprintf(
			`
                MATCH (a:%s)-[r:%s]->(b:%s)
                WHERE a.id = {subjectId} AND b.id = {objectId}
                DELETE r
            `,
			util.UpperCaseFirst(r.Subject), strings.ToUpper(r.Relationship),
			util.UpperCaseFirst(r.Object),
		),
		Parameters: neoism.Props{
			"subjectId": r.SubjectID, "objectId": r.ObjectID,
		},
	}
	return db.Cypher(&cq)
}
