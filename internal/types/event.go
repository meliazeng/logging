package types

type Event struct {
	Channel string
	Value   []byte
}

type PodLabels struct {
	ContainerName string `json:"compute.googleapis.com/resource_name,omitempty"`
	PodId         string `json:"container.googleapis.com/pod_name,omitempty"`
	NameSpaceId   string `json:"container.googleapis.com/namespace_name,omitempty"`
}

type ResourceLabels struct {
	ClusterName   string `json:"cluster_name,omitempty"`
	ContainerName string `json:"container_name,omitempty"`
	InstanceId    string `json:"instance_id,omitempty"`
	NamespaceId   string `json:"namespace_id,omitempty"`
	NamespaceName string `json:"namespace_name,omitempty"`
	PodId         string `json:"pod_id,omitempty"`
	PodName       string `json:"pod_name,omitempty"`
	ProjectId     string `json:"project_id,omitempty"`
	Zone          string `json:"zone,omitempty"`
}

type Resource struct {
	Labels ResourceLabels `json:"labels,omitempty"`
}

type ExtendedEnvelope struct {
	InsertId    string    `json:"insertId,omitempty"`
	Labels      PodLabels `json:"labels,omitempty"`
	Resource    Resource  `json:"resource,omitempty"`
	Severity    string    `json:"severity,omitempty"`
	TextPayload string    `json:"textPayload,omitempty"`
}
