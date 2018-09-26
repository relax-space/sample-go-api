package controllers_test

import (
	"sample-go-api/controllers"
	"testing"

	"github.com/pangpanglabs/goutils/echomiddleware"
)
//go test .\controllers\init_test.go  .\controllers\kafka_consumer_test.go  -v
func Test_Consumer(t *testing.T) {
	controllers.KafkaApiController{}.Consumer(&echomiddleware.KafkaConfig{
		Brokers: []string{"127.0.0.1:9092"},
		Topic:   "kafka_test_topic",
	})
}
