package main

import (
	"github.com/kyeett/restful-diff-detector/webserver"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
)

func TestDiffBasic(t *testing.T) {
	assert.False(t, stringAreEqual(text1, text2))
	assert.True(t, stringAreEqual(text1, text1))
}

func TestDiffHTTP(t *testing.T) {
	handler := http.HandlerFunc(webserver.Users)
	userOne := getStringFromHandler(handler, "/user/1")
	userOneDiff := getStringFromHandler(handler, "/user/1?diff=11")
	userOneAgain := getStringFromHandler(handler, "/user/1")

	// Slightly different texts
	assert.False(t, stringAreEqual(userOne, userOneDiff))

	// Should be same
	assert.True(t, stringAreEqual(userOne, userOneAgain))
}

// TODO, should return err? Need to ask someone :-)
func getStringFromHandler(handler http.Handler, urlString string) string {

	// req, err :=cd . http.NewRequest("GET", , url.Values{"page": {"1"}, "per_page": {"100"}})
	req, err := http.NewRequest("GET", urlString, strings.NewReader("diff=11")) // <-- URL-encoded payload

	if err != nil {
		return "failed" //TODO: learn how to handle error in help methods
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	//    handler := http.HandlerFunc(JSONPage)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		return "failed" //TODO: learn how to handle error in help methods
	}
	return rr.Body.String()
}
