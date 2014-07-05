package booliApi

import (
	"fmt"
	"errors"
	"time"
	"io"
	"crypto/sha1"
	"io/ioutil"
	"math/rand"
	"net/url"
	"net/http"
)

type Client struct {
	Hostname, CallerId, ApiKey string
}

func (this *Client) GetHash (time, unique string) string {
	hash := sha1.New()
	io.WriteString(hash, this.CallerId + time + this.ApiKey + unique)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (this *Client) GetUrl (endpoint string, params map[string]string) string {
	time := GetTime()
	unique := GetUnique()
	hash := this.GetHash(time, unique)

	p := url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}
	p.Set("callerId", this.CallerId)
	p.Set("time", time)
	p.Set("unique", unique)
	p.Set("hash", hash)

	query := p.Encode()

	return "http://" + this.Hostname + "/" + endpoint + "?" + query
}

func (this *Client) Get(endpoint string, params map[string]string) ([]byte, error) {
	u := this.GetUrl(endpoint, params)

	response, err := http.Get(u)
	if(err != nil) {
		return nil, errors.New("Could make GET request to " + u)
	}

	url := this.GetUrl(endpoint, params)

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.booli-v2+json")

	response, err := client.Do(req)

	if(err != nil) {
		return nil, errors.New("Could make GET request to " + url)
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return body, nil
}

func GetTime() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func GetUnique() string {
	random := rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
	return fmt.Sprintf("%d", random)
}
