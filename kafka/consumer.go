package kafka

import (
	"context"
	"github.com/IBM/sarama"
	v1 "github.com/zp857/goutil/constants/v1"
	"github.com/zp857/goutil/threading"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Consumer struct {
	urls        []string
	config      *sarama.Config
	TopicMap    map[string]string
	bindings    []HandlerBinding
	middlewares []MiddlewareFunc
	logger      *zap.SugaredLogger
}

func NewConsumer(urls []string, topicMap map[string]string) (consumer *Consumer) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	return &Consumer{
		urls:        urls,
		config:      config,
		TopicMap:    topicMap,
		bindings:    []HandlerBinding{},
		middlewares: []MiddlewareFunc{},
		logger:      zap.L().Named(v1.KafkaCustomerLogger).Sugar(),
	}
}

// AddMiddleware will add a ServerMiddleware to the list of middlewares to be
func (c *Consumer) AddMiddleware(m MiddlewareFunc) {
	c.middlewares = append(c.middlewares, m)
}

// Bind will add a HandlerBinding to the list of bindings
func (c *Consumer) Bind(bingding HandlerBinding) {
	c.bindings = append(c.bindings, bingding)
}

func (c *Consumer) GetBindings() []HandlerBinding {
	return c.bindings
}

func (c *Consumer) ListenAndServe() {
	c.StartConsume()
}

func (c *Consumer) StartConsume() {
	consumer, err := c.conn()
	if err != nil {
		c.logger.Infof(v1.ConsumerInitError, err)
		return
	}
	for _, binding := range c.bindings {
		go c.consume(consumer, binding)
	}
}

func (c *Consumer) conn() (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(c.urls, c.config)
	return
}

func (c *Consumer) consume(consumer sarama.Consumer, bingding HandlerBinding) {
	// query partitions
	partitions, err := consumer.Partitions(bingding.TopicName)
	if err != nil {
		c.logger.Errorf(v1.GetPartitionsError, err)
		return
	}
	// listen all partitions
	wg := &sync.WaitGroup{}
	for partition := range partitions {
		var destConsumer sarama.PartitionConsumer
		for {
			destConsumer, err = consumer.ConsumePartition(bingding.TopicName, int32(partition), sarama.OffsetNewest)
			if err != nil {
				c.logger.Errorf(v1.ConsumePartitionError, int32(partition), err)
				time.Sleep(defaultSleep)
			} else {
				break
			}
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			for msg := range destConsumer.Messages() {
				// handler
				handler := MiddlewareChain(bingding.HandlerFunc, c.middlewares...)
				ctx := context.Background()
				ctx = context.WithValue(ctx, v1.TopicKey, bingding.TopicName)
				threading.GoSafe(func() {
					handler(ctx, msg.Value)
				})
			}
			defer destConsumer.AsyncClose()
			defer wg.Done()
		}(destConsumer)
	}
	wg.Wait()
	err = consumer.Close()
	if err != nil {
		c.logger.Errorf(v1.ConsumerCloseError, err)
	}
}
