package client

import (
	"fmt"
	"testing"
	"net/url"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestAddsNeededParameters(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(fmt.Sprintf("%s", r.URL))
		fmt.Fprintln(w, u)
	}))

	parameters := map[string]string {
		"limit": "10",
	}

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	response, _ := client.Get("/listing/1234", parameters)

	u, _ := url.Parse(string(response))
	q := u.Query()

	if q.Get("callerId") != "my-caller-id" {
		t.Error("Should have included caller id as parameter")
	}
	if q.Get("limit") != "10" {
		t.Error("Should have included limit as parameter")
	}
	if len(q.Get("unique")) == 0 {
		t.Error("Should have included unique as parameter")
	}
	if len(q.Get("time")) == 0 {
		t.Error("Should have included unique as parameter")
	}
	if len(q.Get("hash")) == 0 {
		t.Error("Should have included hash as parameter")
	}
}

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
		t.Error("Should have set booli id on the fetched listing")
	}

	if err != nil {
		t.Error("Didnt expect an error to have been set")
	}

}

func TestReturnsErrorWhenResponseIsInvalidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "MISSING_SEARCH_PARAMETERS")
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listing, err := client.Listing(1234)

	if listing != nil {
		t.Error("Should have returned a nil pointer to a listing")
	}

	_, wasSyntaxError := err.(*json.SyntaxError)
	if !wasSyntaxError {
		t.Error("Expected a syntax error to have been set")
	}

}

func TestReturnsErrorWhenServerIsNotResponding(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	listing, err := client.Listing(1234)

	if listing != nil {
		t.Error("Should have returned a nil pointer to a listing")
	}

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}
