package sample_test

import (
	"context"
	"os"
	"runtime"

	"github.com/pangpanglabs/goutils/cache"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/echomiddleware"
)

var ctx context.Context
var redisCache *cache.Redis

func init() {
	runtime.GOMAXPROCS(1)
	xormEngine, err := xorm.NewEngine("mysql", os.Getenv("FRUIT_CONN"))
	if err != nil {
		panic(err)
	}
	xormEngine.ShowSQL(true)

	redisCache = cache.NewRedis("redis://127.0.0.1:6379/0")

	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
}
