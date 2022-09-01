package cmd

type NamespaceResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []Namespace `json:"data"`
}

type Namespace struct {
	Namespace         string `json:"namespace"`
	NamespaceShowName string `json:"namespaceShowName"`
	Quota             int    `json:"quota"`
	ConfigCount       int    `json:"configCount"`
	Type              int    `json:"type"`
}

type ServicesResponse struct {
	Count       int       `json:"count"`
	ServiceList []Service `json:"serviceList"`
}

type Service struct {
	ClusterCount         int    `json:"clusterCount"`
	GroupName            string `json:"groupName"`
	HealthyInstanceCount int    `json:"healthyInstanceCount"`
	IpCount              int    `json:"ipCount"`
	Name                 string `json:"name"`
	TriggerFlag          string `json:"triggerFlag"`
}
