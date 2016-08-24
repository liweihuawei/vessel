package kubernetes

import (
	"encoding/json"
	"fmt"
	"log"
)

type ResourceType int

const (
	BINDINGS = iota
	COMPONENTSTATUSES
	ENDPOINTS
	EVENTS
	LIMITRANGES
	NAMESPACES
	NODES
	PERSISTENTVOLUMECLAIMS
	PERSISTENTVOLUMES
	PODS
	PODTEMPLATES
	REPLICATIONCONTROLLERS
	RESOURCEQUOTAS
	SECRETS
	SERVICEACCOUNTS
	SERVICES
	UNKNOWN
)

//TODO more resource type needs to add
func NewResourceType(kind string) ResourceType {
	switch kind {
	case "Binding":
		return BINDINGS
	case "Namespace":
		return NAMESPACES
	case "Node", "NodeLists":
		return NODES
	case "Pod", "PodLists":
		return PODS
	case "ReplicationController":
		return REPLICATIONCONTROLLERS
	case "Service", "ServiceLists":
		return SERVICES
	default:
		return UNKNOWN
	}
}

func (r ResourceType) GetType() string {
	switch r {
	case BINDINGS:
		return "bindings"
	case COMPONENTSTATUSES:
		return "componentstatuses"
	case ENDPOINTS:
		return "endpoints"
	case EVENTS:
		return "events"
	case LIMITRANGES:
		return "limitranges"
	case NAMESPACES:
		return "namespaces"
	case NODES:
		return "nodes"
	case PERSISTENTVOLUMECLAIMS:
		return "persistentvolumeclaims"
	case PERSISTENTVOLUMES:
		return "persistentvolumes"
	case PODS:
		return "pods"
	case PODTEMPLATES:
		return "podtemplates"
	case REPLICATIONCONTROLLERS:
		return "replicationcontrollers"
	case RESOURCEQUOTAS:
		return "resourcequotas"
	case SECRETS:
		return "secrets"
	case SERVICEACCOUNTS:
		return "serviceaccounts"
	case SERVICES:
		return "services"
	default:
		return "unknown"
	}
}

type Params struct {
	ResourceType
	Name         string
	Namespace    string
	IsVisitProxy bool
	IsSetWatcher bool
	Json         []byte
}

type RawParams struct {
	Kind       string         `json:"kind"`
	APIVersion string         `json:"apiVersion"`
	Metadata   ParamsMetaData `json:"metadata"`
}

type ParamsMetaData struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func NewParams(jstr []byte) (*Params, error) {
	var r RawParams
	err := json.Unmarshal(jstr, &r)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if r.APIVersion != "v1" {
		log.Println("Not support version: ", r.APIVersion)
		return nil, err
	}

	return &Params{
		ResourceType: NewResourceType(r.Kind),
		Name:         r.Metadata.Name,
		Namespace:    r.Metadata.Namespace,
		IsVisitProxy: false,
		IsSetWatcher: false,
		Json:         jstr,
	}, nil
}

//TODO add sub path etc.
func (p *Params) BuildPath() (path string) {
	if p.ResourceType == UNKNOWN {
		return
	}

	if p.IsVisitProxy {
		path = "/proxy"
	} else if p.IsSetWatcher {
		path = "/watch"
	}

	if p.Namespace != "" {
		path += fmt.Sprintf("/namespaces/%s/", p.Namespace)
	}

	path += p.GetType()

	if p.Name != "" {
		path += fmt.Sprintf("/%s", p.Name)
	}

	log.Println(path)
	return
}
