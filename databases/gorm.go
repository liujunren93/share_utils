package databases

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Mysql struct {
	Host            string        `json:"host"`
	User            string        `json:"user"`
	Password        string        `json:"password"`
	Port            int           `json:"port"`
	Database        string        `json:"database"`
	ConnMaxLifeTime time.Duration `json:"conn_max_life_time"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	MaxOpenConns    int           `json:"max_open_conns"`
	LogMode         bool          `json:"log_mode"`
}

func (m Mysql) String() string {
	marshal, _ := json.Marshal(&m)
	return string(marshal)
}

// NewMysql will create *gorm.DB
func NewMysql(conf *Mysql) (*gorm.DB, error) {
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai&timeout=5s", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db, err := open.DB()
	if err != nil {
		return nil,err
	}
	db.SetConnMaxLifetime(time.Second * conf.ConnMaxLifeTime)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)
	if conf.LogMode {
		open=open.Debug()
	}
	return open, nil
}
