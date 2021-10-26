module github.com/liujunren93/share_utils

go 1.16

replace (
	github.com/liujunren93/share => ../share
)

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1119
	github.com/aliyun/aliyun-oss-go-sdk v2.1.8+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.2.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.2
	github.com/jinzhu/inflection v1.0.0
	github.com/liujunren93/share v0.0.0-20210821030710-69f159ef04aa
	github.com/micro/go-micro/v2 v2.9.1
	github.com/nacos-group/nacos-sdk-go v0.4.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.7.1
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.231
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v1.0.231
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	go.uber.org/atomic v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	google.golang.org/grpc v1.31.0
	gorm.io/driver/mysql v1.0.2
	gorm.io/gorm v1.21.13
)
