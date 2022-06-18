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
	list = append(list, empData)
	for _, path := range list {
		//fmt.Println(path)
		mapData[path.GetID()] = path
	}

	for _, node := range list {
		if node.GetPid() == node.GetID() {
			continue
		}
		 pid := node.GetPid()
		if _, ok := mapData[pid]; !ok {
			delete(mapData,node.GetID())
			continue
		}
		mapData[pid].AddChild(node)
	}

	return mapData[0]
}
