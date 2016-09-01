package deployment

import (
	"log"
	"strings"

	"github.com/containerops/vessel/models"
	"k8s.io/kubernetes/pkg/api/v1"
)

type Deployment struct {
	ID         int64
	Name       string
	Namespace  string
	DeployType string
	*DeployData
}

func NewDeployment(stage *models.Stage) *Deployment {
	data := &DeployData{}
	replicas := new(int32)
	switch stage.Type {
	case models.STAGECONTAINER:
		*replicas = int32(stage.Replicas)
		data.K8S = &K8SData{
			Name:      stage.Name,
			Namespace: stage.Namespace,
			Replicas:  replicas,
			Labels:    map[string]string{"app": stage.Name},
			PodSpec:   newPodSpec(stage.Artifacts, stage.Volumes),
		}
	case models.STAGEVM:
		return nil
	case models.STAGEPC:
		return nil
	default:
		return nil
	}

	return &Deployment{
		ID:         stage.ID,
		Name:       stage.Name,
		Namespace:  stage.Namespace,
		DeployType: stage.Type,
		DeployData: data,
	}
}

func newPodSpec(a []models.Artifact, v []models.Volume) *v1.PodSpec {
	if len(a) != len(v) {
		log.Println("Length dismatch.")
		return nil
	}

	size := len(a)
	containers := make([]v1.Container, size)
	volumes := make([]v1.Volume, size)
	for i := 0; i < size; i++ {
		containers[i] = v1.Container{
			Name:       a[i].Name,
			Image:      a[i].Path,
			Command:    a[i].Lifecycle.Runtime,
			WorkingDir: a[i].Container.WorkingDir,
			Ports: func(ports []models.ContainerPort) []v1.ContainerPort {
				psize := len(ports)
				parr := make([]v1.ContainerPort, psize)
				for j := 0; j < psize; j++ {
					parr[j] = v1.ContainerPort{
						Name:          ports[j].Name,
						HostPort:      ports[j].HostPort,
						ContainerPort: ports[j].ContainerPort,
					}
				}
				return parr
			}(a[i].Container.Ports),
			Env: func(envs []models.EnvVar) []v1.EnvVar {
				esize := len(envs)
				earr := make([]v1.EnvVar, esize)
				for j := 0; j < esize; j++ {
					earr[j] = v1.EnvVar{
						Name:  envs[j].Name,
						Value: envs[j].Value,
					}
				}
				return earr
			}(a[i].Container.Env),
			Lifecycle: &v1.Lifecycle{
				PostStart: &v1.Handler{
					Exec: &v1.ExecAction{
						Command: a[i].Lifecycle.Before,
					},
				},
				PreStop: &v1.Handler{
					Exec: &v1.ExecAction{
						Command: a[i].Lifecycle.After,
					},
				},
			},
		}

		volumes[i] = v1.Volume{
			Name: v[i].Name,
			VolumeSource: v1.VolumeSource{
				HostPath: &v1.HostPathVolumeSource{
					Path: v[i].HostPath,
				},
			},
		}
	}

	return &v1.PodSpec{
		Volumes:    volumes,
		Containers: containers,
	}
}

func (d *Deployment) Deploy() *models.StageResult {
	switch d.DeployType {
	case models.STAGECONTAINER:
		detail, err := DeployInK8S(d.K8S)
		result := &models.StageResult{
			ID:        d.ID,
			Name:      d.Name,
			Namespace: d.Namespace,
			Detail:    detail,
		}
		if err != nil {
			if strings.Contains(err.Error(), models.ResultTimeout) {
				result.Status = models.ResultTimeout
			} else {
				result.Status = models.ResultFailed
			}
		} else {
			result.Status = models.ResultSuccess
		}
		return result
	case models.STAGEVM:
		return nil
	case models.STAGEPC:
		return nil
	default:
		return nil
	}
}
