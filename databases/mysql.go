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
	DeletedAt uint `sql:"index"`
}

func (b Base) BeforeFind(scope *gorm.Scope) (err error)  {

	scope.SetColumn("DeletedAt","!=0")
	return nil
}
