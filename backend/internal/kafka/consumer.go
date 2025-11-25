package kafka

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	// "time"

	"github.com/segmentio/kafka-go"
)

type ClickEvent struct {
	Slug      string `json:"slug"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Referrer  string `json:"referrer"`
}

func ReadFromKafka() {
    fmt.Println("Kafka consumer init...")

    db, err := sql.Open("postgres", "postgres://user:password@db:5432/urlshortener?sslmode=disable")
    if err != nil {
        log.Fatal("DB connect:", err)
    }

    
    for {
        

        r := kafka.NewReader(kafka.ReaderConfig{
            Brokers:     []string{"kafka:9092"},
            Topic:       "link_clicks",
            Partition:   0,
            MinBytes:    1,
            MaxBytes:    10e6,
            StartOffset: kafka.FirstOffset,
        })



        processKafkaMessages(r, db)


    }
}


func processKafkaMessages(r *kafka.Reader, db *sql.DB) {
	defer r.Close()

	fmt.Println("Inside processKafkaMessages")
	for {
		// fmt.Println("Entered loop")
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Kafka consumer error:", err)
			return // exit this loop → outer loop reconnects
		}

		fmt.Println("Kafka message received:", string(m.Value))

		var event ClickEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		_, err = db.Exec(
			`INSERT INTO url_clicks (slug, count) VALUES ($1, 1)
			ON CONFLICT (slug) DO UPDATE SET count = url_clicks.count + 1`,
			event.Slug,
		)

		if err != nil {
			log.Println("DB update error:", err)
		} else {
			fmt.Printf("Click counted for slug: %s\n", event.Slug)
		}

		_, err = db.Exec(
			`INSERT INTO click_analytics (slug, timestamp, ip, user_agent, referrer)
			VALUES ($1, $2, $3, $4, $5)`,
			event.Slug, event.Timestamp, event.IP, event.UserAgent, event.Referrer,
		)

		if err != nil {
			log.Println("DB analytics insert error:", err)
		}
	}
}
