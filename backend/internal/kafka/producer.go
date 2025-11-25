package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	// "fmt"
	"time"

	"github.com/segmentio/kafka-go"
)



var kafkaWriter *kafka.Writer

func InitKafka() {
	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
    Brokers:      []string{"kafka:9092"},
    Topic:        "link_clicks",
    RequiredAcks: -1,
    Async:        false,
    BatchTimeout: 1 * time.Millisecond,
})
}

func PublishLinkClick(slug, ip, userAgent, referrer string) error {

	payload := map[string]string{
		"slug":      slug,
		"timestamp": time.Now().Format(time.RFC3339),
		"ip":        ip,
		"user_agent": userAgent,
		"referrer":  referrer,
	}
	jsonValue, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	fmt.Println("Publishing Kafka click event for", slug)
	
	msg := kafka.Message{
		Key:   []byte(slug),
		Value: jsonValue,
	}
    return kafkaWriter.WriteMessages(context.Background(), msg)
}
