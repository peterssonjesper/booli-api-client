package client

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestReturnsListingImagesWhenResponseIsValidJson(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"images": [
				{ "id" : 10, "width": 20, "height": 30 }
			]
		}`)
	}))

	type envelope struct {
		Images []map[string]int
	}

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	body, err := client.ListingImages(1234)

	var e envelope
	json.Unmarshal(body, &e)

	if len(e.Images) != 1 {
		t.Error("Expected one image, got %#v", len(e.Images))
	}

	if e.Images[0]["id"] != 10 {
		t.Error("Expected image ID to be 10, was %#v", e.Images[0]["booliId"])
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestReturnsErrorWhenServerIsNotRespondingForListingImages(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	_, err := client.ListingImages(1234)

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}
