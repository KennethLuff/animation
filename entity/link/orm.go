package link

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"animation/entity"
)

var Orm *xorm.Engine

func Init(dns string) error {
	var err error
	Orm, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		return err
	}

	Orm.ShowSQL(true)
	err = Orm.Sync2(new(entity.User))

	if err != nil {
		fmt.Println("data source init error", err.Error())
	}
	return err
}