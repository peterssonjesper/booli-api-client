package client

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func getListing(testServer *httptest.Server, id int) ([]byte, error) {
	client := New(testServer.URL, "my-caller-id", "my-api-key")
	return client.Listing(id)
}

func TestReturnsListingWhenResponseIsValidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"listings": [
				{ "booliId" : 1234 }
			]
		}`)
	}))

	type envelope struct {
		Listings []map[string]int
	}

	body, err := getListing(testServer, 1234)

	var e envelope
	json.Unmarshal(body, &e)

	if e.Listings[0]["booliId"] != 1234 {
		t.Error("Expected booli ID to be 1234, was %#v", e.Listings[0]["booliId"])
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestReturnsErrorWhenServerIsNotRespondingForListing(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	_, err := getListing(testServer, 1234)

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}

func TestCallsCorrectUrlWhenFetchingListing(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL)
	}))

	url, _ := getSimilarListings(testServer, 1234)

	expected := "/listings/1234"
	if string(url)[:len(expected)] != expected {
		t.Errorf("Expected url to start with %s, was %s", expected, url)
	}
}
