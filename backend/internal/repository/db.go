package repository

import (
	"database/sql"
	"errors"
	"time"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitPostgres(dsn string) {
	var err error

	fmt.Println("Connecting to Postgres with DSN: ok ", dsn)


	// Retry logic for DB startup (optional, but good for Docker Compose setups)
	for range [10]int{} {
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
		expires_at TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create urls table: %v", err)
	}
	createTableQuery = `
	CREATE TABLE IF NOT EXISTS url_clicks (
		slug TEXT PRIMARY KEY,
		count INTEGER NOT NULL DEFAULT 0
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create url_clicks table: %v", err)
	}
	createTableQuery = `
	CREATE TABLE IF NOT EXISTS click_analytics (
		id SERIAL PRIMARY KEY,
		slug TEXT,
		timestamp TIMESTAMP,
		ip TEXT,
		user_agent TEXT,
		referrer TEXT
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create click_analytics table: %v", err)
	}
	

	log.Println("Database initialized and urls table ensured.")
}

// Auto-create table if not exists
	// createTableQuery := `CREATE TABLE IF NOT EXISTS urls (
	// 	id SERIAL PRIMARY KEY,
	// 	slug TEXT UNIQUE NOT NULL,
	// 	long_url TEXT NOT NULL,
	// 	clicks INTEGER DEFAULT 0,
		
	// );`

	// _, err = db.Exec(createTableQuery)

func Save(slug string, longURL string, expiresAt *time.Time) error {
	_, err := db.Exec(`INSERT INTO urls (slug, long_url, expires_at) VALUES ($1, $2, $3)`, slug, longURL, expiresAt)
	if err != nil {
		log.Println("Error in Save: ", err)
	}
	return err
}

func Find(slug string) (string, *time.Time,  error) {
	var longURL string
	var expiresAt *time.Time
	err := db.QueryRow(`SELECT long_url, expires_at FROM urls WHERE slug = $1`, slug).Scan(&longURL, &expiresAt)
	// if(err != nil){
	// 	fmt.Println("Looking for slug:", slug)
	// 	fmt.Println("Looking for string:", longURL)
	// 	fmt.Println("to this is error", err)
	// }
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil, errors.New("URL not found2")
		}
		return "", nil, err
	}
	return longURL, expiresAt, nil
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

func GetClickCount(slug string) (int, *time.Time, error) {
	click := 0
	var expiresAt *time.Time
	err := db.QueryRow(`SELECT clicks, expires_at FROM urls WHERE slug = $1`, slug).Scan(&click, &expiresAt)
	if err != nil {
		return 0,expiresAt, err
	}
	return click,expiresAt, nil
}


func DB() *sql.DB {
    return db
}
