package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"sample-go-api/factory"

	"github.com/pangpanglabs/goutils/echomiddleware"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaApiController struct {
}

func (d KafkaApiController) Init(g *echo.Group) {
	g.GET("/producers", d.Producer)
}

func (d KafkaApiController) Producer(c echo.Context) error {
	var producer *kafka.Producer
	config := factory.ConfigKafka(c.Request().Context(), "sample_kafka")

	if p, err := kafka.NewProducer(config.Brokers, config.Topic, func(c *sarama.Config) {
		c.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
		c.Producer.Compression = sarama.CompressionGZIP     // Compress messages
		c.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	}); err != nil {
		logrus.Error("Create Kafka Producer Error", err)
	} else {
		producer = p
	}
	var dto = TestDto{
		"xiao",
		18,
	}
	producer.Send(&dto)
	fmt.Println("kafka producer")

	return c.String(http.StatusOK, "success")
}

func (d KafkaApiController) Consumer(k *echomiddleware.KafkaConfig) {
	fmt.Println("====>1")
	consumer, err := kafka.NewConsumer(k.Brokers, k.Topic, nil, sarama.OffsetNewest, func(c *sarama.Config) {

	})
	if err != nil {
		panic(err)
	}
	fmt.Println("====>2")
	messages, err := consumer.Messages()

	if err != nil {
		panic(err)
	}
	fmt.Println("====>3")
	for m := range messages {
		var v TestDto
		fmt.Println("====>4")
		d := json.NewDecoder(bytes.NewReader(m.Value))
		d.UseNumber()
		err := d.Decode(&v)
		if err != nil {
			//log.Println(err)
			fmt.Println(err)
		}
		fmt.Println("kafka consumered", v)
		//fmt.Printf("consumer=>%+v", v)
	}
}

type TestDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
