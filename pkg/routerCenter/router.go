package routerCenter

import (
	"context"
)

type Router struct {
	Method      string                 `json:"method" yaml:"method"` // http mehtod
	GrpcMenthod string                 `json:"grpc_method" yaml:"grpc_method"`
	ReqParams   map[string]interface{} `json:"req_params" yaml:"req_params"`
}

type WathMethod string

const (
	REGISTRY WathMethod = "reg"
	DEL      WathMethod = "del"
)

type RouterCentry struct {
	Namespace string `json:"namespace"`
	Prefix    string
}

func (r *RouterCentry) GetKey(app string) string {
	if r.Prefix == "" {
		r.Prefix = "routerCenter"
	}
	if r.Namespace == "" {
		r.Namespace = "default"
	}
	key := r.Prefix + "/" + r.Namespace + "/"
	if app != "" {
		key += app
	}
	return key
}

func (r *RouterCentry) GetKeys(app string) string {

	return r.GetKey(app) + "*"
}

type RouterCenter interface {
	GetAllRouter(ctx context.Context) map[string]map[string]*Router        //map[app]map[apiPath]Router
	GetRouter(ctx context.Context, app string) (map[string]*Router, error) //map[apiPath]map[]
	Registry(ctx context.Context, app string, router map[string]*Router) error
	DelRouter(ctx context.Context, app string) error
	Watch(ctx context.Context, callback func(app string, router map[string]*Router, err error))
}
