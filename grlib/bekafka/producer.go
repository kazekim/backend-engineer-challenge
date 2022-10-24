package bekafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"log"
)

type Producer struct {
	conn sarama.SyncProducer
}

// NewProducer init kafka producer
func (c *defaultClient) NewProducer() (*Producer, grerrors.Error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	log.Println(c.cfg.BrokerUrl)
	conn, err := sarama.NewSyncProducer(c.cfg.BrokerUrl, config)
	if err != nil {
		return nil, grerrors.NewKafkaProducerError(err)
	}

	return &Producer{
		conn: conn,
	}, nil
}

// Close close kafka producer
func (p *Producer) Close() grerrors.Error {
	err := p.conn.Close()
	if err != nil {
		return grerrors.NewDefaultError(err)
	}
	return nil
}

// SendMessage start sending message to kafka with topic
func (p *Producer) SendMessage(topic, message string) grerrors.Error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := p.conn.SendMessage(msg)
	if err != nil {
		return grerrors.NewKafkaProducerError(err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
