package models_test

import (
	"context"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/echomiddleware"
)

var ctx context.Context
var db *xorm.Engine

func init() {
	runtime.GOMAXPROCS(1)
	var err error
	db, err = xorm.NewEngine("mysql", os.Getenv("FRUIT_CONN"))
	if err != nil {
		panic(err)
	}
	db.ShowSQL(true)

	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, db.NewSession())
}
