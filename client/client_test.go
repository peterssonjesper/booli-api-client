package client

import (
	"fmt"
	"testing"
	"net/url"
	"net/http"
	"net/http/httptest"
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
