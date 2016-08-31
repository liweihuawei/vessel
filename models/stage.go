package models

import (
	"time"
)

//Stage Type
const (
	STAGECONTAINER = "container"
	STAGEVM        = "vm"
	STAGEPC        = "pc"
)

// Stage stage data
type Stage struct {
	ID            int64      `json:"-"`
	Name          string     `json:"name"  binding:"Required" gorm:"varchar(30)"`
	Namespace     string     `json:"namespace"  binding:"Required" gorm:"varchar(30)"`
	Type          string     `json:"type"  binding:"In(container,vm,pc)" gorm:"varchar(20)"`
	Replicas      uint       `json:"replicas" gorm:"type:int;default:0"`
	Dependencies  string     `json:"dependencies,omitempty" gorm:"varchar(255)"` // eg : "a,b,c"
	Artifacts     []Artifact `json:"artifacts" binding:"Required" gorm:"-"`
	Volumes       []Volume   `json:"volumes,omitempty" gorm:"-"`
	ArtifactsJSON string     `json:"-" gorm:"column:artifacts;type:text;not null"` // json type
	VolumesJSON   string     `json:"-" gorm:"column:volumes;type:text;not null"`   // json type
	Status        uint       `json:"status" gorm:"type:int;default:0"`
	Created       *time.Time `json:"created" `
	Updated       *time.Time `json:"updated"`
	Deleted       *time.Time `json:"deleted"`
}

// StageVersion data
type StageVersion struct {
	ID      int64      `json:"id" gorm:"primary_key"`
	PvID    int64      `json:"pvid" gorm:"type:int;not null"`
	SID     int64      `json:"sid" gorm:"type:int;not null"`
	Detail  string     `json:"detail" gorm:"type:text;"`
	Status  uint       `json:"status" gorm:"type:int;default:0"`
	Created *time.Time `json:"created" `
	Updated *time.Time `json:"updated"`
	Deleted *time.Time `json:"deleted"`
	Stage   Stage      `json:"-"`
}

// Artifact data
type Artifact struct {
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Lifecycle *Lifecycle `json:"lifecycle,omitempty"`
	Container *Container `json:"container,omitempty"`
}

// Lifecycle data
type Lifecycle struct {
	Before  []string `json:"before,omitempty"`
	Runtime []string `json:"runtime,omitempty"`
	After   []string `json:"after,omitempty"`
}

// Container data
type Container struct {
	WorkingDir string          `json:"workingDir,omitempty"`
	Ports      []ContainerPort `json:"ports,omitempty"`
	Env        []EnvVar        `json:"env,omitempty"`
}

// ContainerPort data
type ContainerPort struct {
	Name          string `json:"name,omitempty"`
	HostPort      int32  `json:"hostPort,omitempty"`
	ContainerPort int32  `json:"containerPort,omitempty"`
}

// EnvVar data
type EnvVar struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// Volume data
type Volume struct {
	Name     string `json:"name,omitempty"`
	HostPath string `json:"hostPath,omitempty"`
}

// StageResult stage result
type StageResult struct {
	Namespace string `json:"-"`
	ID        int64  `json:"stageVersionID"`
	Name      string `json:"stageName"`
	Status    string `json:"runResult"`
	Detail    string `json:"detail"`
}
