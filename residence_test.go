package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnsResidencesWhenResponseIsValidJson(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"residences": [
				{ "id" : 1234 }
			]
		}`)
	}))

	type envelope struct {
		Residences []map[string]int
	}

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	body, err := client.Residences()

	var e envelope
	json.Unmarshal(body, &e)

	if len(e.Residences) != 1 {
		t.Error("Expected one sold property, got %#v", len(e.Residences))
	}

	if e.Residences[0]["id"] != 1234 {
		t.Error("Expected booli ID to be 1234, was %#v", e.Residences[0]["id"])
	}

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

}

func TestJSONEncodesEstimatePayload(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, fmt.Sprintf("%s", body))
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	body, err := client.Estimate(map[string]string{
		"hello": "world",
	})

	if err != nil {
		t.Error("Expected error to be nil, was %#v", err)
	}

	expected := `{"hello":"world"}`
	if string(body) != expected {
		t.Errorf("Expected body to be encoded to %s, but was %s", expected, body)
	}

}
