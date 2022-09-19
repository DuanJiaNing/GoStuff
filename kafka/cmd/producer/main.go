package main

import (
	"context"
	"flag"
	"fmt"
	"kafka/pkg/log"
	"kafka/pkg/otel"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

var (
	brokers = flag.String("brokers", "localhost:9092", "The Kafka brokers to connect to, as a comma separated list")
	topic   = flag.String("topic", "default_topic", "The Kafka topic to use")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	log.Init("[producer]")
	shutdownTraceProvider, err := otel.InstallFilePipeline(ctx, "producer")
	if err != nil {
		log.L.Fatalf("failed to initialize stdout export pipeline: %v", err)
	}
	defer shutdownTraceProvider()

	if *brokers == "" {
		log.L.Fatalln("at least one broker is required")
	}
	splitBrokers := strings.Split(*brokers, ",")
	sarama.Logger = log.L

	// simple sarama producer that adds a new producer interceptor
	conf := sarama.NewConfig()
	conf.Version = sarama.V0_11_0_0
	conf.Producer.Interceptors = []sarama.ProducerInterceptor{NewOTelInterceptor(splitBrokers)}
	producer, err := sarama.NewAsyncProducer(splitBrokers, conf)
	if err != nil {
		panic("Couldn't create a Kafka producer")
	}
	defer producer.AsyncClose()

	// kill -2, trap SIGINT to trigger a shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// ticker
	bulkSize := 2
	duration := 5 * time.Second
	ticker := time.NewTicker(duration)
	log.L.Printf("Starting to produce %v messages every %v", bulkSize, duration)
	for {
		select {
		case t := <-ticker.C:
			now := t.Format(time.RFC3339)
			log.L.Printf("\nproducing %v messages to topic %s at %s", bulkSize, *topic, now)
			for i := 0; i < bulkSize; i++ {
				producer.Input() <- &sarama.ProducerMessage{
					Topic: *topic, Key: nil,
					Value: sarama.StringEncoder(fmt.Sprintf("testmessage_%v/%v: %s", i+1, bulkSize, now)),
				}
			}
		case <-signals:
			log.L.Println("terminating the program")
			log.L.Println("Bye :)")
			return
		}
	}
}
