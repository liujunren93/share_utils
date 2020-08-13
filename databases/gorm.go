package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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




// NewMysql will create *gorm.DB
func NewMysql(conf *Mysql) (*gorm.DB, error) {
	open, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai", conf.User, conf.Password, conf.Host, conf.Port, conf.Database))
	if err != nil {
		return nil, err
	}
	open.DB().SetConnMaxLifetime(time.Second * conf.ConnMaxLifeTime)
	open.DB().SetMaxOpenConns(conf.MaxOpenConns)
	open.DB().SetMaxOpenConns(conf.MaxOpenConns)
	open.LogMode(conf.LogMode)
	return open, nil
}
