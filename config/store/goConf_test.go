package store

import (
	"encoding/json"
	"fmt"
	"github.com/shareChina/utils/config"
	"github.com/shareChina/utils/log"
	"testing"
)


//
type Mysql struct {
	Host     string `yml:"host"`
	User     string `yml:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}


func TestNewGoConf(t *testing.T) {


}

func init() {
	//getConf(BaseConf.Acm.AccessKey, BaseConf.Acm.SecretKey, BaseConf.App.ServerName, BaseConf.Acm.Endpoint)
	getConf("LTAI4G58UvgfoChctGeeiZTS", "S62Ik50uEKbjLERA3fGa0ZIKVoJWf9", "da1af185-ef8b-4fd5-ab5a-1c15fc3fe906", "acm.aliyun.com")
}
var MysqlConf *Mysql
////获取mysql
func getConf(accessKey, secretKey, namespaceID, endpoint string) {

	newStore, err :=NewAliyunStore(accessKey, secretKey, namespaceID, endpoint)
	if err != nil {
		log.Logger.Fatal(err)
	}
	err = config.GetConfig(newStore, &MysqlConf, "mysql", "test")
	if err != nil {
		log.Logger.Fatal(err,"111")
	}
	fmt.Println(MysqlConf,err)
	config.ListenConfig(newStore, mysqlConfigUpdate, "mysql",  "test")
}
func mysqlConfigUpdate(data interface{}) {
	json.Unmarshal([]byte(data.(string)), &MysqlConf)
}
