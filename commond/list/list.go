package list

/**
* @Author: liujunren
* @Date: 2022/4/5 15:54
 */

type TreeNoder interface {
	GetID() uint
	GetPid() uint
	GetChilds() interface{}
	AddChild(interface{})
}

func List2Tree(list []TreeNoder, empData TreeNoder) TreeNoder {
	var mapData = make(map[uint]TreeNoder)
	mapData[0] = empData
	for _, path := range list {
		mapData[path.GetID()] = path
	}
	for _, node := range list {
		if node.GetPid() == node.GetID() {
			continue
		}
		mapData[node.GetPid()].AddChild(node)
	}

	return mapData[0]
}
