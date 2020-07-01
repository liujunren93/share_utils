package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// NewGormDB will create *gorm.DB
func NewGormDB(user, passowrd, host, database string, port int) (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai", user, passowrd, host, port, database))
}

type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type ModelError struct {
	Code int //500 系统异常 501 sql异常，以及其他
	Msg  interface{}
}

func (Base) NewError(code int, err interface{}) *ModelError {
	return &ModelError{Code: code, Msg: err}
}
