package router

import (
	"path"
	"strings"
)

type Node struct {
	Path     string
	GrpcPath string
	Method   string
	Childs   []*Node
}

func NewTree(root, grpcPath string) *Node {
	return &Node{Path: root, GrpcPath: grpcPath}
}
func (node *Node) Add(reqPath string, method, grpcPath string) {
	method = strings.ToUpper(method)
	reqPath = path.Clean(reqPath)
	reqPath = strings.Trim(reqPath, "/")
	node.add(strings.Split(reqPath, "/"), method, grpcPath)
}

func (node *Node) add(paths []string, method, grpcPath string) {
	var isNew = true
	var tmpNode *Node
	var cpath = paths[0]
	for _, v := range node.Childs {
		if v.Path == cpath {
			if len(paths) == 1 {
				if v.Method == method {
					tmpNode = v
					tmpNode.GrpcPath = grpcPath
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
		tmpNode.add(paths[1:], method, grpcPath)
	}

	if len(paths) == 1 {
		tmpNode.GrpcPath = grpcPath
		tmpNode.Method = method

	}

	if isNew {
		node.Childs = append(node.Childs, tmpNode)
	}

}

func (node *Node) Find(reqPath string, method string) *Node {
	method = strings.ToUpper(method)
	reqPath = "/" + reqPath
	reqPath = path.Clean(reqPath)
	paths := strings.Split(reqPath, "/")
	paths[0] = "/"
	return node.find(paths, method)
}

func (node *Node) find(paths []string, method string) *Node {
	var cpath = paths[0]
	if node.Path == cpath {
		if len(paths) == 1 && node.Method == method {
			return node
		} else {
			for _, n := range node.Childs {
				if len(paths) > 1 {
					tnode := n.find(paths[1:], method)
					if tnode != nil {
						return tnode
					}
				}

			}
		}
	} else {
		if len(paths) == 1 && string(node.Path[0]) == ":" && node.Method == method {
			return node
		}

	}
	return nil
}
