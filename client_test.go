package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetAddsNeededParameters(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(fmt.Sprintf("%s", r.URL))
		fmt.Fprintln(w, u)
	}))

	parameters := map[string]string{
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

func TestPostAddsNeededParameters(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _ := url.Parse(fmt.Sprintf("%s", r.URL))
		fmt.Fprintln(w, u)
	}))

	parameters := map[string]string{
		"limit": "10",
	}

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	response, _ := client.Post("/listing/1234", []byte(""), parameters)

	u, _ := url.Parse(string(response))
	q := u.Query()

	if q.Get("callerId") != "my-caller-id" {
		t.Error("Should have included caller id as parameter")
	}
}

func TestPostAddsPayload(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, fmt.Sprintf("%s", body))
	}))

	client := New(testServer.URL, "my-caller-id", "my-api-key")
	response, _ := client.Post("/listing/1234", []byte("this-is-the-payload"))

	if string(response) != "this-is-the-payload" {
		t.Error("Should have included the payload in the POST call")
	}
}
