package kubernetes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"k8s.io/kubernetes/pkg/api/v1"
)

const (
	HOST = "http://10.67.147.217:8080"
)

var client RESTClient

func init() {
	client = NewRESTClient(HOST)
}

func TestResource(t *testing.T) {
	data, err := ioutil.ReadFile("./testpod.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	params, err := NewParams(data)
	if err != nil {
		t.Error(err)
		return
	}

	result := client.Create(params)
	log.Println(result)

	result = client.Get(params)
	log.Println(result)

	result = client.Delete(params)
	log.Println(result)
}
