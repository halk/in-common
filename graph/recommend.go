package graph

import (
	"fmt"
	"github.com/halk/in-common/model"
	"github.com/halk/in-common/util"
	"github.com/jmcvetta/neoism"
	"strings"
)

// GetRecommendations returns a list of recommended IDs
func GetRecommendations(rq *model.RecommendationRequest) (*model.Recommendation, error) {
	// connect to graph
	db := GetGraph()

	res := []struct {
		ID     string  `json:"object.id"`
		Weight float32 `json:"weight"`
	}{}

	cq := neoism.CypherQuery{
		Statement: fmt.Sprintf(
			// Currently we are using only COUNT
			// Since we store the number of times the relational event happened
			// We could have used SUM(r), however that is not useful for the
			// current use cases (a customer may view a product many times and
			// skew the results)
			`
				MATCH (subject:%[1]s)-[:%[2]s]->(%[3]s)<-[:%[2]s]-(%[1]s)-[r:%[2]s]->(object:%[3]s)
				WHERE subject.id = {subjectId} AND NOT((subject)-[:%[2]s]->(object))
				RETURN object.id, COUNT(r) AS weight
				ORDER BY weight DESC
				LIMIT %[4]d
		  `,
			util.UpperCaseFirst(rq.Subject), strings.ToUpper(rq.Relationship),
			util.UpperCaseFirst(rq.Object), rq.Limit,
		),
		Parameters: neoism.Props{"subjectId": rq.SubjectID},
		Result:     &res,
	}
	if err := db.Cypher(&cq); err != nil {
		return nil, err
	}

	recommendation := model.NewRecommendation()
	for _, result := range res {
		recommendation.Add(result.ID, result.Weight)
	}
	return recommendation, nil
}
