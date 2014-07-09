package client

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestReturnsSimilarListingWhenResponseIsValidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"listings": [
				{ "booliId" : 1234 }
			]
		}`)
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listings, err := client.SimilarListings(1234)

	if len(listings) != 1 {
		t.Error("Expected one listing, got %#v", len(listings))
	}

	if listings[0].BooliId != 1234 {
		t.Error("Expected booli ID to be 1234, was %#v", listings[0].BooliId)
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestReturnsErrorWhenResponseForSimilarListingsIsInvalidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "MISSING_SEARCH_PARAMETERS")
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listings, err := client.SimilarListings(1234)

	if listings != nil {
		t.Error("Expected listings to be nil, was %#v", listings)
	}

	_, wasSyntaxError := err.(*json.SyntaxError)
	if !wasSyntaxError {
		t.Error("Expected a syntax error to have been set")
	}

}

func TestReturnsErrorWhenServerIsNotRespondingForSimilarListings(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listings, err := client.SimilarListings(1234)

	if listings != nil {
		t.Error("Expected listings to be nil, was %#v", listings)
	}

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}
