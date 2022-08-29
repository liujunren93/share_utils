package router

import "context"

type Router struct {
	Method    string                 `json:"method" yaml:"method"`
	ReqParams map[string]interface{} `json:"req_params" yaml:"req_params"`
}

type RouterCentry struct {
	Namespace string `json:"namespace"`
}

func (r *RouterCentry) GetKey(key string) string {
	return "routerCenter/" + r.Namespace + "/" + key
}

type RouterCenter interface {
	GetRouter(ctx context.Context, key string) map[string]Router //map[method:path]map[]
	Registry(ctx context.Context, key string, router map[string]Router) error
	DelRouter(ctx context.Context, key string) error
	Watch(ctx context.Context, key string, callback func(router map[string]Router, error error))
}
