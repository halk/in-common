package graph

import (
	"fmt"
	"github.com/jmcvetta/neoism"
	"github.com/stretchr/testify/assert"
	"inCommon/model"
	"inCommon/util"
	"strings"
	"testing"
)

func TestGetGraph(t *testing.T) {
	db := GetGraph()
	assert.IsType(t, neoism.Database{}, *db, "Unexpected type")
}

func TestGetOrCreateNode(t *testing.T) {
	db := GetGraph()

	node1, created1 := createNode(db, "graphtest", "node1", t)
	if node1 != nil {
		assert.True(t, created1, "Unexpected created flag value")
		assert.Equal(t, "node1", node1.Data["id"], "Unexptected value for ID")
	}

	node2, created2 := createNode(db, "graphtest", "node1", t)
	if node2 != nil {
		assert.False(t, created2, "Unexpected created flag value")
		assert.Equal(t, "node1", node2.Data["id"], "Unexptected value for ID")

		if node1 != nil {
			assert.Equal(t, node2.Id(), node1.Id(), "Unexpected mismatch of Neo4j IDs")
		}
		deleteNode(db, node2)
	}
}

func TestCreateOrIncRelationship(t *testing.T) {
	db := GetGraph()

	node1, _ := createNode(db, "graphtest2_1", "node1", t)
	node2, _ := createNode(db, "graphtest2_2", "node2", t)

	if node1 != nil && node2 != nil {
		r := model.Relationship{"graphtest2_1", "node1", "graphtest2_2", "node2", "tested"}
		if err := CreateOrIncRelationship(db, &r); err != nil {
			assert.Fail(t, "Unexpected error: "+err.Error())
		}
		freq := getFrequencyRelations(db, &r, t)
		if freq > -1 {
			assert.Equal(t, 1, freq, "Unexpected value for frequency (times the same relationship happened)")
		}

		if err := CreateOrIncRelationship(db, &r); err != nil {
			assert.Fail(t, "Unexpected error: "+err.Error())
		}
		freq = getFrequencyRelations(db, &r, t)
		if freq > -1 {
			assert.Equal(t, 2, freq, "Unexpected value for frequency (times the same relationship happened)")
		}

		deleteNode(db, node1)
		deleteNode(db, node2)
	}
}

func TestDeleteRelationship(t *testing.T) {
	db := GetGraph()

	node1, _ := createNode(db, "graphtest2_1", "node1", t)
	node2, _ := createNode(db, "graphtest2_2", "node2", t)

	if node1 != nil && node2 != nil {
		r := model.Relationship{"graphtest2_1", "node1", "graphtest2_2", "node2", "tested"}
		if err := CreateOrIncRelationship(db, &r); err != nil {
			assert.Fail(t, "Unexpected error: "+err.Error())
		}
		if err := DeleteRelationship(db, &r); err != nil {
			assert.Fail(t, "Unexpected error: "+err.Error())
		}
	} else {
		assert.Fail(t, "Could not create nodes to test delete relationship")
	}
}

func deleteNode(db *neoism.Database, node *neoism.Node) {
	// need to remove all relationships before
	rels, err := node.Relationships()
	if err != nil {
		panic(err)
	}
	for _, rel := range rels {
		rel.Delete()
	}
	node.Delete()
}

func reloadNode(db *neoism.Database, label string, id string) {
	// neoism does not provide find, simpler to createOrGet than writing a CypherQuery
	node, _, err := GetOrCreateNode(db, label, "id", neoism.Props{"id": id})
	if err != nil {
		return
	}
	if err := node.Delete(); err != nil {
		panic(err)
	}
}

func createNode(db *neoism.Database, label string, id string, t *testing.T) (*neoism.Node, bool) {
	node, created, err := GetOrCreateNode(db, label, "id", neoism.Props{"id": id})
	if err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
		return nil, false
	}
	return node, created
}

func getFrequencyRelations(db *neoism.Database, r *model.Relationship, t *testing.T) int {
	res := []struct {
		Frequency int `json:"frequency"`
	}{}

	cq := neoism.CypherQuery{
		Statement: fmt.Sprintf(
			`
            	MATCH (subject:%[1]s)-[r:%[2]s]->(object:%[3]s)
            	WHERE subject.id = {subjectId} AND object.id = {objectId}
            	RETURN r.frequency as frequency
            `,
			util.UpperCaseFirst(r.Subject), strings.ToUpper(r.Relationship),
			util.UpperCaseFirst(r.Object),
		),
		Parameters: neoism.Props{"subjectId": r.SubjectID, "objectId": r.ObjectID},
		Result:     &res,
	}
	if err := db.Cypher(&cq); err != nil {
		assert.Fail(t, "Unexpected error: "+err.Error())
		return -1
	}

	if len(res) > 0 {
		return res[0].Frequency
	}
	return -1
}
