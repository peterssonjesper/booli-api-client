package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnsSoldPropertyWhenResponseIsValidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"sold": [
				{ "booliId" : 1234 }
			]
		}`)
	}))

	type envelope struct {
		Sold []map[string]int
	}

	body, err := getSoldProperty(testServer, 1234)

	var e envelope
	json.Unmarshal(body, &e)

	if e.Sold[0]["booliId"] != 1234 {
		t.Error("Expected booli ID to be 1234, was %#v", e.Sold[0]["booliId"])
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestReturnsErrorWhenServerIsNotRespondingForSoldProperty(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	_, err := getSoldProperty(testServer, 1234)

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}

func TestCallsCorrectUrlWhenFetchingSoldProperty(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL)
	}))

	url, _ := getSoldProperty(testServer, 1234)

	expected := "/sold/1234"
	if string(url)[:len(expected)] != expected {
		t.Errorf("Expected url to start with %s, was %s", expected, url)
	}
}

func getSoldProperty(testServer *httptest.Server, id int) ([]byte, error) {
	client := New(testServer.URL, "my-caller-id", "my-api-key")
	return client.SoldProperty(id)
}
