package entity

type Registry struct {
	Enable bool   `yaml:"enable" json:"enable"`
	Type   string `json:"type" yaml:"type"`
	Etcd   *Etcd  `json:"etcd" yaml:"etcd"`
}

type Etcd struct {
	Endpoints []string `json:"endpoints" yaml:"endpoints"`
}
