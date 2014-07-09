package client

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestReturnsListingWhenResponseIsValidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"listings": [
				{ "booliId" : 1234 }
			]
		}`)
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listing, err := client.Listing(1234)

	if listing.BooliId != 1234 {
		t.Error("Expected booli ID to be 1234, was %#v", listing.BooliId)
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestReturnsErrorWhenResponseForListingIsInvalidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "MISSING_SEARCH_PARAMETERS")
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listing, err := client.Listing(1234)

	if listing != nil {
		t.Error("Expected listing to be nil, was %#v", listing)
	}

	_, wasSyntaxError := err.(*json.SyntaxError)
	if !wasSyntaxError {
		t.Error("Expected a syntax error to have been set")
	}

}

func TestReturnsErrorWhenServerIsNotRespondingForListing(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listing, err := client.Listing(1234)

	if listing != nil {
		t.Error("Expected listing to be nil, was %#v", listing)
	}

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}
