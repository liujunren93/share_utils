module github.com/liujunren93/share_utils

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/liujunren93/share => ../share

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1119 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.1.8+incompatible // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.2.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.2
	github.com/jinzhu/inflection v1.0.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/liujunren93/share v0.0.0-20210123125154-38f91d2b1a1e
	github.com/micro/go-micro/v2 v2.9.1
	github.com/nacos-group/nacos-sdk-go v0.4.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	google.golang.org/grpc v1.31.0
	gopkg.in/ini.v1 v1.62.0 // indirect
	gorm.io/driver/mysql v1.0.2
	gorm.io/gorm v1.20.2
)
