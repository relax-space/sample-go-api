package sample_test

import (
	"testing"
	"time"

	"github.com/pangpanglabs/goutils/test"
)

const (
	PROJECTNAME = "sample"
)

func TestRedisStore(t *testing.T) {
	var v interface{}
	_, err := redisCache.LoadOrStore(PROJECTNAME+":key", &v, func() (interface{}, error) {
		return "value", nil
	})
	test.Ok(t, err)
	test.Equals(t, v, "value")
	time.Sleep(10 * time.Minute)
}

func TestRedisLoad(t *testing.T) {
	var v interface{}
	loadFromCache, err := redisCache.LoadOrStore(PROJECTNAME+":key", &v, nil)
	test.Equals(t, true, loadFromCache)
	test.Ok(t, err)
	test.Equals(t, v, "value")
	time.Sleep(10 * time.Minute)
}
