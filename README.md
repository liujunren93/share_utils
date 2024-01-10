# share_utils app公共工具
## 配置文件
### local config.yaml 
``` yaml
app_name: "example-app" // app name
namespace: "shareLife" // 命名空间
version: "1.0" // 版本号
run_mode: "debug"
listen_addr: "" // 监听端口
conf_center: // 配置中心 目前仅支持redis
  enable: true
  type: "redis"
  conf_name: "example-app"
  group: "debug" // 配置分组
  config: // 配置中心配置
    mode: 1  // general:1 cluster:2 sentinel:3
    network: "tcp"
    addr: "node1:6379"
    username: "1"
    password: "1"
    // mode=2 cluster
    mode: 2  
    network: "tcp"
    addrs: 
      - "node1:6379"
      - "node2:6379"
    username: "1"
    password: "1"

```
### cloud config.json 配置中心
```json
{
	"version": "1",
	"redis": { // 同local
		"mode": 1,
		"network": "tcp",
		"addr": "redis:6379"
	},
	"registry": { // 注册中心
		"type": "etcd",// etcd k8s 
		// k8s 只需要设置type 
		"etcd": {
			"endpoints": [
				"etcd0:2379",
				"etcd1:2379",
				"etcd2:2379"
			]
		}
	},
	"router_center": {// 自动化路由注册 目前只支持redis   注册key "router/" + namespace + "/" + router_prefix + "/"
		"type": 1,
		"enable": true,
		"app_prefix": "share_app", // gin 路由前缀
		"router_prefix": "share_api_client",  // 路由注册中心前缀
		"redis": {
			"mode": 1,
			"network": "tcp",
			"addr": "redis:6379"
		}
	},
	"log": {
		"out": "console",
		"set_report_caller": true,
		"level": "debug",
		"rotate": {
			"log_file": "./log/log.log",
			"max_age": 86400,
			"rotation": 6000
		}
	}

}

``` 