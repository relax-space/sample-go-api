package controllers

import (
	"context"
	"fmt"
	"sample-go-api/factory"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pangpanglabs/goutils/kafka"
)

//If you need to call the producer multiple times in a loop,
// in order to save server resources, it is recommended to make a single case.
var kafkaSampleInstance *KafkaSample
var kafkaSampleOnce sync.Once

type KafkaSample struct {
	Producer *kafka.Producer
}

func (KafkaSample) GetInstance(ctx context.Context) *KafkaSample {
	kafkaSampleOnce.Do(func() {
		config := factory.ConfigKafka(ctx, "sample_kafka")
		if p, errd := kafka.NewProducer(config.Brokers, config.Topic, func(c *sarama.Config) {
			c.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
			c.Producer.Compression = sarama.CompressionGZIP     // Compress messages
			c.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

		}); errd != nil {
			// to be
			fmt.Println(errd)
			return
		} else {
			kafkaSampleInstance = &KafkaSample{
				Producer: p,
			}
		}
	})
	return kafkaSampleInstance
}
