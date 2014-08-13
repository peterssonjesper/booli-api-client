package client

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func getListingAreas(testServer *httptest.Server, id int) ([]byte, error) {
	client := New(testServer.URL, "my-caller-id", "my-api-key")
	return client.ListingAreas(id)
}

func TestReturnsListingAreasWhenResponseIsValidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"areas": [
				{ "id" : 1234 }
			]
		}`)
	}))

	type envelope struct {
		Areas []map[string]int
	}

	body, err := getListingAreas(testServer, 1234)

	var e envelope
	json.Unmarshal(body, &e)

	if len(e.Areas) != 1 {
		t.Error("Expected one sold property, got %#v", len(e.Areas))
	}

	if e.Areas[0]["id"] != 1234 {
		t.Error("Expected booli ID to be 1234, was %#v", e.Areas[0]["booliId"])
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestReturnsErrorWhenServerIsNotRespondingWithListingAreas(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	_, err := getListingAreas(testServer, 1234)

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}

func TestCallsCorrectUrlWhenFetchingListingAreas(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL)
	}))

	url, _ := getListingAreas(testServer, 1234)

	expected := "/listings/1234/areas"
	if string(url)[:len(expected)] != expected {
		t.Errorf("Expected url to start with %s, was %s", expected, url)
	}
}
