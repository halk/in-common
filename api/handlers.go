// The inCommon/api package serves the API
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/schema"
	"inCommon/engine"
	"inCommon/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the inCommon recommendation engine!")
}

// AddEvent handler
func AddEvent(w http.ResponseWriter, r *http.Request) {
	// decode json / unmarshal to Relationship model
	var relationship model.Relationship
	if err := json.Unmarshal(readBody(r), &relationship); err != nil {
		returnError(w, 422, err)
		return
	}

	if err := verifyRelationship(&relationship); err != nil {
		returnError(w, 422, err)
		return
	}

	// process event
	if err := engine.ProcessAddEvent(&relationship); err != nil {
		returnError(w, 422, err)
		return
	}

	// send response
	w.WriteHeader(http.StatusNoContent)
}

// RemoveEvent handler
func RemoveEvent(w http.ResponseWriter, r *http.Request) {
	// decode json / unmarshal to Relationship model
	var relationship model.Relationship

	if err := json.Unmarshal(readBody(r), &relationship); err != nil {
		returnError(w, 422, err)
		return
	}

	if err := verifyRelationship(&relationship); err != nil {
		returnError(w, 422, err)
		return
	}

	// process event
	if err := engine.ProcessRemoveEvent(&relationship); err != nil {
		returnError(w, 422, err)
		return
	}

	// send response
	w.WriteHeader(http.StatusNoContent)
}

// Recommend handler
func Recommend(w http.ResponseWriter, r *http.Request) {
	// get variables
	rq := model.RecommendationRequest{}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	if err := decoder.Decode(&rq, r.URL.Query()); err != nil {
		returnError(w, http.StatusInternalServerError, err)
		return
	}

	// get recommendations
	recommendations, err := engine.GetRecommendations(&rq)
	if err != nil {
		returnError(w, http.StatusInternalServerError, err)
		return
	}

	setJSONEncoding(w)
	w.WriteHeader(http.StatusOK)
	encodeJSON(w, recommendations.Results)
}

// util methods

func readBody(r *http.Request) []byte {
	// read body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 262144))
	if err != nil {
		panic(err)
	}
	// close body
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	return body
}

func verifyRelationship(r *model.Relationship) error {
	if r.Subject == "" || r.SubjectID == "" || r.ObjectID == "" ||
		r.Object == "" || r.Relationship == "" {
		return errors.New("Incomplete relationship data")
	}
	return nil
}

func setJSONEncoding(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func encodeJSON(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func returnError(w http.ResponseWriter, status int, err error) {
	log.Printf("Error: %s", err)
	setJSONEncoding(w)
	w.WriteHeader(422)
	encodeJSON(w, err.Error())
}
