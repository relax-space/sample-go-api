package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"sample-go-api/factory"
	"sample-go-api/models"

	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/httpreq"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/kafka"
	"github.com/sirupsen/logrus"
)

type KafkaApiController struct {
}

func (d KafkaApiController) Init(g *echo.Group) {
	g.GET("/producers", d.Producer)
	g.GET("/producers/fruits", d.ProducerFruit)
}

func (d KafkaApiController) Producer(c echo.Context) error {

	var dto = TestDto{
		"xiao",
		18,
	}
	KafkaSample{}.GetInstance(c.Request().Context()).Producer.Send(&dto)
	fmt.Println("kafka producer")

	return c.String(http.StatusOK, "success")
}

func (d KafkaApiController) Consumer(k *echomiddleware.KafkaConfig) {
	consumer, err := kafka.NewConsumer(k.Brokers, k.Topic, nil, sarama.OffsetNewest, func(c *sarama.Config) {

	})
	if err != nil {
		panic(err)
	}
	messages, err := consumer.Messages()

	if err != nil {
		panic(err)
	}
	for m := range messages {
		var v TestDto
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

func (d KafkaApiController) ProducerFruit(c echo.Context) error {
	var producer *kafka.Producer
	config := factory.ConfigKafka(c.Request().Context(), "fruit_kafka")

	if p, err := kafka.NewProducer(config.Brokers, config.Topic, func(c *sarama.Config) {
		c.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
		c.Producer.Compression = sarama.CompressionGZIP     // Compress messages
		c.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	}); err != nil {
		logrus.Error("Create Kafka Producer Error", err)
	} else {
		producer = p
	}
	url := fmt.Sprintf("%v/fruits", factory.ConfigString(c.Request().Context(), "sample_url"))
	var respDto struct {
		Result struct {
			Items []models.FruitLog `json:"items"`
		} `json:"result"`
		Success bool     `json:"success"`
		Error   ApiError `json:"error"`
	}
	_, err := httpreq.New(http.MethodGet, url, nil).Call(&respDto)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorServiceUnavailable, err)
	}
	for k := range respDto.Result.Items {
		respDto.Result.Items[k].Success = true
	}
	producer.Send(&respDto.Result.Items)
	fmt.Println("kafka producer:", respDto.Result)

	return c.String(http.StatusOK, "success")
}

func (d KafkaApiController) ConsumerFruit(db *xorm.Engine, k *echomiddleware.KafkaConfig) {
	consumer, err := kafka.NewConsumer(k.Brokers, k.Topic, nil, sarama.OffsetNewest, func(c *sarama.Config) {

	})
	if err != nil {
		panic(err)
	}
	messages, err := consumer.Messages()

	if err != nil {
		panic(err)
	}
	for m := range messages {
		v := make([]*models.FruitLog, 0)
		d := json.NewDecoder(bytes.NewReader(m.Value))
		d.UseNumber()
		err := d.Decode(&v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("kafka consumerFruit")
		affectedRow, err := (models.FruitLog{}).CreateBatch(db, v)
		fmt.Println(fmt.Sprintf("affectedRow:%v, err:%v", affectedRow, err))
	}
}
