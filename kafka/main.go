package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ngaut/log"
	"sync"
	"time"
)

func main() {
	SaramaConsumer()
}

func SaramaProducer() {
	fmt.Printf("consumer_test\n")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// producer
	producer, err := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	defer producer.AsyncClose()

	var wg sync.WaitGroup
	go func(p sarama.AsyncProducer) {
		wg.Add(1)
		for{
			select {
			case  <-p.Successes():
				fmt.Println("xxxxxx")
				// fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
		wg.Done()
	}(producer)
	message := &sarama.ProducerMessage{
		Topic:     "test1",
		Key:       sarama.ByteEncoder("xxxxxxxxxxxxxxxxx"),
		Value:     sarama.ByteEncoder("xsadfasdfasdfasdf123123123213"),
		Timestamp: time.Time{},
	}
	producer.Input() <- message
	wg.Wait()
}

func SaramaConsumer() {

	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()
	var wg sync.WaitGroup

	for i := int32(0); i < 10; i++ {
		partitionConsumer, err := consumer.ConsumePartition("test1", i, sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("try create partitionConsumer error %s\n", err.Error())
			return
		}
		wg.Add(1)
		go func(partitionConsumer sarama.PartitionConsumer) {
			for j := 0; j <1; j++ {
				msg := <-partitionConsumer.Messages()
				fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
			}
			wg.Done()
			if err := partitionConsumer.Close(); err != nil {
				log.Error(err)
			}
		}(partitionConsumer)
	}
	wg.Wait()
}
