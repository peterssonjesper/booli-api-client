package client

import (
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestReturnsPreviousSalesWhenResponseIsValidJson(t *testing.T) {

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

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	body, err := client.PreviousSales(1234)

	var e envelope
	json.Unmarshal(body, &e)

	if len(e.Sold) != 1 {
		t.Error("Expected one sold property, got %#v", len(e.Sold))
	}

	if e.Sold[0]["booliId"] != 1234 {
		t.Error("Expected booli ID to be 1234, was %#v", e.Sold[0]["booliId"])
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestReturnsErrorWhenServerIsNotRespondingWithPreviousSales(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, `{"error": "Internal server error"}`)
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	_, err := client.PreviousSales(1234)

	if err == nil {
		t.Error("Expected an error to have been set")
	}

}
