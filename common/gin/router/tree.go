package router

import (
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liujunren93/share_utils/helper"
)

type Node struct {
	Path                 string
	GrpcPath             string
	Method               string
	MiddlewaresWhitelist []string
	ReqParams            map[string]interface{} `json:"req_params" yaml:"req_params"`
	Childs               []*Node
}

func NewTree(root, grpcPath string) *Node {
	return &Node{Path: root, GrpcPath: grpcPath}
}
func (node *Node) Add(reqPath string, method, grpcPath string, mWhitelist []string, reqParams map[string]interface{}) {
	method = strings.ToUpper(method)
	reqPath = path.Clean(reqPath)
	reqPath = strings.Trim(reqPath, "/")
	node.add(strings.Split(reqPath, "/"), method, grpcPath, mWhitelist, reqParams)
}

func (node *Node) add(paths []string, method, grpcPath string, mWhitelist []string, reqParams map[string]interface{}) {
	var isNew = true
	var tmpNode *Node
	var cpath = paths[0]
	for _, v := range node.Childs {
		if v.Path == cpath {
			if len(paths) == 1 {
				if v.Method == method {
					tmpNode = v
					tmpNode.GrpcPath = grpcPath
					tmpNode.ReqParams = reqParams
					isNew = false
					break
				}
			} else {
				tmpNode = v
				isNew = false
				break
			}
		}
	}
	if tmpNode == nil {
		tmpNode = new(Node)
		tmpNode.Path = cpath
	}

	if len(paths) > 1 {
		tmpNode.add(paths[1:], method, grpcPath, mWhitelist, reqParams)
	}

	if len(paths) == 1 {
		tmpNode.GrpcPath = grpcPath
		tmpNode.Method = method
		tmpNode.ReqParams = reqParams
		tmpNode.MiddlewaresWhitelist = mWhitelist
	}

	if isNew {
		node.Childs = append(node.Childs, tmpNode)
	}

}

func (node *Node) Find(reqPath string, method string) (n *Node, params gin.Param) {
	method = strings.ToUpper(method)
	reqPath = "/" + reqPath
	reqPath = path.Clean(reqPath)
	paths := strings.Split(reqPath, "/")
	paths[0] = "/"
	return node.find(paths, method)
}

func (node *Node) find(paths []string, method string) (n *Node, param gin.Param) {
	var cpath = paths[0]
	if node.Path == cpath {
		if len(paths) == 1 && node.Method == method {
			return node, param
		} else {
			for _, n := range node.Childs {
				if len(paths) > 1 {
					tnode, param := n.find(paths[1:], method)
					if tnode != nil {
						return tnode, param
					}
				}

			}
		}
	} else {
		if len(paths) == 1 && string(node.Path[0]) == ":" && node.Method == method {
			return node, gin.Param{
				Key:   helper.SubstrRight(node.Path, ":"),
				Value: cpath,
			}
		}

	}
	return
}
