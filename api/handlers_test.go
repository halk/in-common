package api

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)

	m.Run()
	deleteEvent("tester1", "test1", nil)
	deleteEvent("tester2", "test1", nil)
	deleteEvent("tester2", "test2", nil)
	deleteEvent("tester2", "test3", nil)
	deleteEvent("tester3", "test2", nil)
}

func TestIndex(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := match(r, t)

	if w != nil {
		assert.Equal(
			t,
			"Welcome to the inCommon recommendation engine!",
			w.Body.String(),
			"Unexpected response for Index handler",
		)
	}
}

func TestAddEvent(t *testing.T) {
	w := postEvent("tester1", "test1", t)

	if w != nil {
		assert.Equal(t, 204, w.Code, "Unexpected HTTP status code in response")
	}
}

func TestAddEvent_IncompleteRelationshipError(t *testing.T) {
	body := `{"subject":"tester"}`
	r, _ := http.NewRequest("POST", "/event", bytes.NewBufferString(body))
	w := match(r, t)

	if w != nil {
		assert.Equal(t, 422, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "\"Incomplete relationship data\"\n", w.Body.String(), "Incorrect error message in response")
	}
}

func TestAddEvent_UnmarshalingError(t *testing.T) {
	body := `{"subject":tester"}`
	r, _ := http.NewRequest("POST", "/event", bytes.NewBufferString(body))
	w := match(r, t)

	if w != nil {
		assert.Equal(t, 422, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "\"invalid character 'e' in literal true (expecting 'r')\"\n", w.Body.String(), "Incorrect error message in response")
	}
}

func TestRemoveEvent(t *testing.T) {
	w := deleteEvent("tester1", "test1", t)

	if w != nil {
		assert.Equal(t, 204, w.Code, "Unexpected HTTP status code in response")
	}
}

func TestRemoveEvent_IncompleteRelationshipError(t *testing.T) {
	body := `{"subject":"tester"}`
	r, _ := http.NewRequest("DELETE", "/event", bytes.NewBufferString(body))
	w := match(r, t)

	if w != nil {
		assert.Equal(t, 422, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "\"Incomplete relationship data\"\n", w.Body.String(), "Incorrect error message in response")
	}
}

func TestRemoveEvent_UnmarshalingError(t *testing.T) {
	body := `{"subject":tester"}`
	r, _ := http.NewRequest("DELETE", "/event", bytes.NewBufferString(body))
	w := match(r, t)

	if w != nil {
		assert.Equal(t, 422, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "\"invalid character 'e' in literal true (expecting 'r')\"\n", w.Body.String(), "Incorrect error message in response")
	}
}

func TestRecommend_Empty(t *testing.T) {
	w := recommend("tester1", 10, t)
	if w != nil {
		assert.Equal(t, 200, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "{}\n", w.Body.String(), "Unexpected response")
	}
}

func TestRecommend(t *testing.T) {
	postEvent("tester1", "test1", t)
	postEvent("tester2", "test1", t)
	postEvent("tester2", "test2", t)

	w := recommend("tester1", 10, t)
	if w != nil {
		assert.Equal(t, 200, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "{\"test2\":1}\n", w.Body.String(), "Unexpected response")
	}
}

func TestRecommend_Limit(t *testing.T) {
	postEvent("tester2", "test3", t)
	postEvent("tester3", "test1", t)
	postEvent("tester3", "test2", t)

	w := recommend("tester1", 10, t)
	if w != nil {
		assert.Equal(t, 200, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "{\"test2\":2,\"test3\":1}\n", w.Body.String(), "Unexpected response")
	}

	w = recommend("tester1", 1, t)
	if w != nil {
		assert.Equal(t, 200, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "{\"test2\":2}\n", w.Body.String(), "Unexpected response")
	}

	deleteEvent("tester2", "test2", t)
	deleteEvent("tester3", "test2", t)
	w = recommend("tester1", 1, t)
	if w != nil {
		assert.Equal(t, 200, w.Code, "Unexpected HTTP status code in response")
		assert.Equal(t, "{\"test3\":1}\n", w.Body.String(), "Unexpected response")
	}
}

// utility methods

func match(r *http.Request, t *testing.T) *httptest.ResponseRecorder {
	router := NewRouter()
	match := new(mux.RouteMatch)
	if matched := router.Match(r, match); !matched {
		if t != nil {
			assert.Fail(t, "Request could not be matched")
		} else {
			errors.New("Request could not be matched")
		}
		return nil
	}

	w := httptest.NewRecorder()
	match.Handler.ServeHTTP(w, r)
	return w
}

func postEvent(subjectId string, objectId string, t *testing.T) *httptest.ResponseRecorder {
	return sendEvent("POST", subjectId, objectId, t)
}

func deleteEvent(subjectId string, objectId string, t *testing.T) *httptest.ResponseRecorder {
	return sendEvent("DELETE", subjectId, objectId, t)
}

func sendEvent(method string, subjectId string, objectId string, t *testing.T) *httptest.ResponseRecorder {
	body := fmt.Sprintf(
		`{"subject":"tester","subject_id":"%s","object":"test","object_id":"%s","relationship":"tested"}`,
		subjectId, objectId,
	)
	r, _ := http.NewRequest(method, "/event", bytes.NewBufferString(body))
	return match(r, t)
}

func recommend(subjectId string, limit int, t *testing.T) *httptest.ResponseRecorder {
	r, _ := http.NewRequest("GET", fmt.Sprintf(
		"/recommend?subject=tester&subject_id=%s&object=test&relationship=tested&limit=%d",
		subjectId, limit,
	), nil)
	return match(r, t)
}
