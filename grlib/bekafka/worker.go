package bekafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"os"
	"os/signal"
	"syscall"
)

type Worker struct {
	conn sarama.Consumer
}

//NewWorker init new kafka worker
func (c *defaultClient) NewWorker() (*Worker, grerrors.Error) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(c.cfg.BrokerUrl, config)
	if err != nil {
		return nil, grerrors.NewKafkaConsumerError(err)
	}

	return &Worker{
		conn: conn,
	}, nil
}

//ConsumePartition set waiting to receive kafka message from producer
func (w *Worker) ConsumePartition(topic string, action func(topic, message string) WorkerStatus) grerrors.Error {

	consumer, err := w.conn.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return grerrors.NewKafkaConsumerError(err)
	}

	fmt.Printf("Worker %v is started\n", topic)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	doneCh := make(chan WorkerStatus)
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				topic := msg.Topic
				message := string(msg.Value)
				fmt.Printf("Received message : | Topic(%s) | Message(%s) \n", string(msg.Topic), string(msg.Value))
				status := action(topic, message)
				if status == WorkerStatusTerminate {
					doneCh <- status
				}
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- WorkerStatusTerminate
			}
		}
	}()

	<-doneCh

	if err := w.conn.Close(); err != nil {
		return grerrors.NewKafkaConsumerError(err)
	}
	return nil
}
