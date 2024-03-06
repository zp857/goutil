package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/dustin/go-humanize"
	"github.com/zp857/goutil/constants/v1"
	"github.com/zp857/goutil/jsonutil"
	"github.com/zp857/goutil/slice"
	"go.uber.org/zap"
)

type Producer struct {
	urls          []string
	config        *sarama.Config
	logger        *zap.SugaredLogger
	debugPrint    bool
	ignoredTopics []string
}

func NewProducer(urls []string, debugPrint bool, ignoreTopics []string) *Producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	return &Producer{
		urls:          urls,
		config:        config,
		logger:        zap.L().Named(v1.KafkaProducerLogger).Sugar(),
		debugPrint:    debugPrint,
		ignoredTopics: ignoreTopics,
	}
}

func (p *Producer) SendJSON(topic string, obj any) {
	producer, err := sarama.NewSyncProducer(p.urls, p.config)
	if err != nil {
		p.logger.Errorf(v1.RequestError, err)
		return
	}
	defer producer.Close()
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(obj)
	if err != nil {
		p.logger.Errorf(v1.RequestError, err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonBytes),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		p.logger.Errorf(v1.RequestError, err)
		return
	}
	if !p.debugPrint && slice.Contain(p.ignoredTopics, topic) {
		return
	}
	p.logger.Infof(v1.ProducerRequest, topic, partition, offset, humanize.Bytes(uint64(msg.Value.Length())), jsonutil.MustPretty(obj))
}
