package gorm

import (
	"sync"

	"gorm.io/gorm"
)

var DB *gorm.DB
var mu sync.Mutex
var dbVersion int64

func InitGorm(conf *Mysql) error {
	return newDB(conf)
}

func newDB(conf *Mysql) (err error) {
	tmpVersion := dbVersion
	mu.Lock()
	defer mu.Unlock()
	if tmpVersion == dbVersion {
		DB, err = NewMysql(conf, nil)
	}

	return err
}

func UpdateDB(conf *Mysql) error {
	return newDB(conf)
}
