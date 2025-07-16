package repository

import (
	"database/sql"
	"errors"
	"time"

	// "fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitPostgres(dsn string) {
	var err error

	// Retry logic for DB startup (optional, but good for Docker Compose setups)
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("Waiting for database... (retrying in 1s)\n")
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create table if not exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		slug TEXT UNIQUE NOT NULL,
		long_url TEXT NOT NULL,
		clicks INTEGER DEFAULT 0
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create urls table: %v", err)
	}

	log.Println("Database initialized and urls table ensured.")
}

// Auto-create table if not exists
	// createTableQuery := `
	// // CREATE TABLE IF NOT EXISTS urls (
	// // 	id SERIAL PRIMARY KEY,
	// // 	slug TEXT UNIQUE NOT NULL,
	// // 	long_url TEXT NOT NULL,
	// // 	clicks INTEGER DEFAULT 0,
		
	// // );`

	// _, err = db.Exec(createTableQuery)

func Save(slug string, longURL string) error {
	_, err := db.Exec(`INSERT INTO urls (slug, long_url) VALUES ($1, $2)`, slug, longURL)
	if err != nil {
		log.Println("Error in Save: ", err)
	}
	return err
}

func Find(slug string) (string, error) {
	var longURL string
	err := db.QueryRow(`SELECT long_url FROM urls WHERE slug = $1`, slug).Scan(&longURL)
	// if(err != nil){
	// 	fmt.Println("Looking for slug:", slug)
	// 	fmt.Println("Looking for string:", longURL)
	// 	fmt.Println("to this is error", err)
	// }
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("URL not found2")
		}
		return "", err
	}
	return longURL, nil
}

func Exists(slug string) (bool, error) {
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM urls WHERE slug = $1)`,slug).Scan(&exists)
	return exists, err
}


func IncrementClicks(slug string) error {
	_, err := db.Exec(`UPDATE urls SET clicks = clicks + 1 WHERE slug = $1`, slug)
	return err
}

func GetClickCount(slug string) (int, error) {
	click := 0
	err := db.QueryRow(`SELECT clicks FROM urls WHERE slug = $1`, slug).Scan(&click)
	if err != nil {
		return 0,err
	}
	return click, nil
}
