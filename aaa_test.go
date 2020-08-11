package utils

import (
	"fmt"
	"github.com/shareChina/utils/config"
	"github.com/shareChina/utils/config/store"
	"net/http"
	"testing"
)

func TestDeleteConfig(t *testing.T) {

	var endpoint = "acm.aliyun.com:8080"
	var namespaceId = "da1af185-ef8b-4fd5-ab5a-1c15fc3fe906"
	var accessKey = "LTAI4G58UvgfoChctGeeiZTS"
	var secretKey = "S62Ik50uEKbjLERA3fGa0ZIKVoJWf9"

	acmStore, _ := store.NewAcmStore(accessKey, secretKey, namespaceId, endpoint, "", "")
	//acmStore.ListenConfig(func(i interface{}) {
	//	fmt.Println(i)
	//},"mysql","test")
	config.ListenConfig(acmStore, func(i interface{}) {
		fmt.Println(i)
	},"mysql","test")
	http.ListenAndServe(":11111",nil)

}
