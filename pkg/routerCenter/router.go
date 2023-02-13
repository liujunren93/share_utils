package routerCenter

import (
	"context"
	"fmt"
	"reflect"

	"github.com/liujunren93/share_utils/helper"
)

type Binding string

const (
	JSON Binding = "json"
	FORM Binding = "form"
)

type Router struct {
	GrpcMenthod          string                 `json:"grpc_method" yaml:"grpc_method"`
	Codes                Binding                `json:"codes" yaml:"codes"`
	ReqParams            map[string]interface{} `json:"req_params" yaml:"req_params"`
	MiddlewaresWhitelist []string               `json:"middl_white_list" yaml:"middl_white_list"`
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
	key := "router/" + r.Namespace + "/" + r.Prefix + "/"

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
	Lease(ctx context.Context, app string) error
	DelRouter(ctx context.Context, app string) error
	Watch(ctx context.Context, callback func(app string, router map[string]*Router, err error))
}

func BuildRouter(grpcMethod string, req interface{}, middlWhitelist ...string) *Router {
	var router = Router{
		GrpcMenthod:          grpcMethod,
		MiddlewaresWhitelist: middlWhitelist,
		ReqParams:            map[string]interface{}{},
	}
	var reqParams = make(map[string]interface{})
	t := reflect.TypeOf(req)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if t.Field(i).Name[0] >= 97 {
			continue
		}
		if t.Field(i).Tag.Get("json") == "-" {
			continue
		}
		validate := t.Field(i).Tag.Get("binding")
		if validate != "" {
			reqParams[helper.SnakeString(t.Field(i).Name)] = validate
		}

	}
	router.ReqParams = reqParams
	return &router
}
