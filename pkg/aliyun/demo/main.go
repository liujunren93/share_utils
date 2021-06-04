package main

import (
	"encoding/json"
	"fmt"
	"github.com/liujunren93/share_utils/pkg/aliyun/shsts"
	"net/http"
)

func main() {
	http.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			fmt.Println(origin)
			w.Header().Set("Access-Control-Allow-Origin", "http://demo.com:8001")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers",
				"Action, Module,authorization")   //有使用自定义头 需要这个,Action, Module是例子
		}
		if r.Method == "OPTIONS" {
			return
		}
		sts, err := shsts.NewSTS("LTAI5tRVGzhVvY49HfUotYkw", "UDbhrwCp6xMRCNlqk5tpLgA21lnwiC")
		if err != nil {
			panic(err)
		}
		res:=make(map[string]interface{})
		credentials, _, err := sts.Credentials("acs:ram::1855340179767769:role/osssts", "")
		if err != nil {
			panic(err)
		}
		res["code"]=200
		res["data"]= map[string]interface{}{"dir":"rim","bucketName":"share-life-public","credentials":credentials,"roleArn":"acs:ram::1855340179767769:role/osssts"}
		marshal, _ := json.Marshal(res)
		w.Write(marshal)

	})


	http.ListenAndServe(":9999",nil)
}
