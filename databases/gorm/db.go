package gorm

import (
	"sync"

	"gorm.io/gorm"
)

var DB *gorm.DB
var mu sync.Mutex

func InitGorm(conf *Mysql) error {
	return newDB(conf)
}

func newDB(conf *Mysql) (err error) {
	mu.Lock()
	defer mu.Unlock()
	DB, err = NewMysql(conf, nil)
	return err
}

func UpdateDB(conf *Mysql) error {
	return newDB(conf)
}
