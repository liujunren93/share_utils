package routerCenter

import "context"

type Router struct {
	Method      string                 `json:"method" yaml:"method"` // http mehtod
	GrpcMenthod string                 `json:"grpc_method" yaml:"grpc_method"`
	ReqParams   map[string]interface{} `json:"req_params" yaml:"req_params"`
}

type RouterCentry struct {
	Namespace string `json:"namespace"`
}

func (r *RouterCentry) GetKey(app string) string {
	return "routerCenter/" + r.Namespace + "/" + app
}

type RouterCenter interface {
	GetAllRouter(ctx context.Context) (map[string]map[string]Router, error) //map[app]map[apiPath]Router
	GetRouter(ctx context.Context, app string) (map[string]Router, error)   //map[apiPath]map[]
	Registry(ctx context.Context, app string, router map[string]Router) error
	DelRouter(ctx context.Context, app string) error
	Watch(ctx context.Context, app string, callback func(router map[string]Router, error error))
}
