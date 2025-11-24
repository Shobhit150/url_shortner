package kafka

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"os/signal"

	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	// "github.com/Shobhit150/url_shortner/internal/kafka"
)

type ClickEvent struct {
	Slug      string `json:"slug"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Referrer  string `json:"referrer"`
}

func ReadFromKafka() {
	fmt.Println(" Done okok ")
	db, err := sql.Open("postgres", "postgres://user:password@db:5432/urlshortener?sslmode=disable")

	if err != nil {
		log.Fatal("DB connect:", err)
	}
	defer db.Close()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "link_clicks",
		GroupID: "click_analytics",
	})
	defer r.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	fmt.Println("Kafka consumer started. Listening for click events...")

loop:
	for {
		select {
		case <-sigChan:
			break loop
		default:
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("ReadMessage error:", err)
				continue
			}
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
}
