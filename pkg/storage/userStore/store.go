package userStore

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/liujunren93/share_utils/auth"
	"github.com/liujunren93/share_utils/auth/jwt"
	"github.com/liujunren93/share_utils/log"
	"time"
)

const (
	storeAgentPrefix       = "rbacUserLoginData"
	storePermissionPrefix  = "rbacUserPermission"
	OperatingPermissionKey = "operatingCompanyPermission" //运营权限
	RootPermissionKey      = "rootCompanyPermission"      //root权限
)

type UserStore struct {
	Expire int64  // 保持登录时间
	Secret string // 提取key密钥
	Redis  *redis.Client
}

func NewUserStore(expire int64, secret string, redis *redis.Client) *UserStore {
	return &UserStore{
		Expire: expire,
		Redis:  redis,
		Secret: secret,
	}
}

func (s *UserStore) getKey(c *gin.Context) (string, bool) {
	token := c.GetHeader("Authorization")
	auth := jwt.NewAuth(auth.WithSecret(s.Secret))
	inspect, err := auth.Inspect(token)

	if err != nil {
		return "", false
	}
	if claims, ok := inspect.(*jwt.JwtClaims); ok {
		s := claims.Data.(string)
		return s, true
	}
	return "", false
}

//Store 存储登录信息
func (s *UserStore) Store(key string, l *LoginInfo) error {
	ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
	l.CreateAt = time.Now().Local().Unix()
	infoStr, _ := encode(l)
	set := s.Redis.Set(ctxTimeout, storeAgentPrefix+key, infoStr, time.Duration(s.Expire)*time.Second)
	return set.Err()
}

// StorePermission 存储权限
// role =1  root,role=2 运营
func (s *UserStore) StorePermission(info *LoginInfo, isRoot bool, permissions []Permission) error {
	marshal, err := json.Marshal(&permissions)
	if err != nil {
		return err
	}
	ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
	if info == nil { // 机构
		if isRoot { //root
			return s.Redis.Set(ctxTimeout, RootPermissionKey, string(marshal), 0).Err()
		} else {
			return s.Redis.Set(ctxTimeout, OperatingPermissionKey, string(marshal), 0).Err()
		}

	} else { //后台管理
		return s.Redis.Set(ctxTimeout, fmt.Sprintf("%s_%d", storePermissionPrefix, info.UID), string(marshal), time.Duration(s.Expire)*time.Second).Err()
	}
}

// LoadPermission 获取权限
// role =1  root,role=2 运营
func (s *UserStore) LoadPermission(info *LoginInfo) ([]Permission, error) {
	ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
	var data []Permission
	var key string
	var expire time.Duration
	if info.UserType == 2 { // 机构
		if info.IsRoot { //root
			key = RootPermissionKey
		} else {
			key = OperatingPermissionKey
		}

	} else {
		expire = time.Duration(s.Expire) * time.Second
		key = fmt.Sprintf("%s_%d", storePermissionPrefix, info.UID)
	}
	get := s.Redis.Get(ctxTimeout, key)
	if permission, err := get.Bytes(); err == nil {
		go func() { // 续命
			if info.UserType!=2 {
				ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
				s.Redis.Expire(ctxTimeout, storeAgentPrefix+key, expire)
			}

		}()
		return data, json.Unmarshal(permission, &data)
	} else {
		return nil, err
	}

}

//Store 存储登录信息
func (s *UserStore) LoadByKey(key string) (*LoginInfo, bool) {
	ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
	get := s.Redis.Get(ctxTimeout, storeAgentPrefix+key)
	if get.Err() != nil {
		return nil, false
	}
	bytes, err := get.Bytes()
	if err != nil {
		log.Logger.Error(err)
		return nil, false
	}

	info, err := decode(bytes)
	go func() { // 续命
		ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
		s.Redis.Expire(ctxTimeout, storeAgentPrefix+key, time.Duration(s.Expire)*time.Second)
	}()
	return info, true
}

//Load 获取用户登录信息
func (s *UserStore) Load(c *gin.Context) (*LoginInfo, bool) {
	key, ok := s.getKey(c)
	if !ok {
		return nil, false
	}
	return s.LoadByKey(key)

}

//Count 在线用户统计
func (s *UserStore) Count() int {
	ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
	keys := s.Redis.Keys(ctxTimeout, storeAgentPrefix+"*")
	return len(keys.Val())
}

//Del
func (s *UserStore) Del(key string) {
	go func() {
		ctxTimeout, _ := context.WithTimeout(context.TODO(), time.Second*3)
		err := s.Redis.Del(ctxTimeout, storeAgentPrefix+key).Err()
		log.Logger.Error(err)
	}()
}
