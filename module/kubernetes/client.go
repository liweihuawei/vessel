package kubernetes

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	_APIPATH = "/api/v1"
)

type RESTClient struct {
	base             *url.URL
	versionedAPIPath string
	Client           *http.Client
}

func NewRESTClient(host string) *RESTClient {
	return &RESTClient{
		base: &url.URL{
			Host: host,
		},
		versionedAPIPath: _APIPATH,
		Client:           &http.Client{},
	}
}

func (c *RESTClient) Get(params *Params) string {
	path := params.BuildPath()
	if path == "" {
		return "Not supported yet"
	}

	path = c.base.Host + c.versionedAPIPath + path
	log.Println(path)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err.Error()
	}

	return c.doRequest(req)
}

func (c *RESTClient) Create(params *Params) string {
	if params.Json == nil {
		return "Empty request body."
	}

	path := params.BuildPathForPost()
	if path == "" {
		return "Not supported yet"
	}

	path = c.base.Host + c.versionedAPIPath + path
	log.Println(path)
	req, err := http.NewRequest("POST", path, bytes.NewReader(params.Json))
	if err != nil {
		return err.Error()
	}

	return c.doRequest(req)
}

func (c *RESTClient) Delete(params *Params) string {
	path := params.BuildPath()
	if path == "" {
		return "Not supported yet"
	}

	path = c.base.Host + c.versionedAPIPath + path
	log.Println(path)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return err.Error()
	}

	return c.doRequest(req)
}

func (c *RESTClient) doRequest(req *http.Request) string {
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
