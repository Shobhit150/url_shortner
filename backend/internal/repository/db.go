package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitPostgres(dsn string) {
	var err error

	

	// Retry logic for DB startup (good inside Docker)
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

	// ------------------------------------------------------------------------------------
	// TABLE DEFINITIONS
	// ------------------------------------------------------------------------------------

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		slug TEXT UNIQUE NOT NULL,
		long_url TEXT NOT NULL,
		expires_at TIMESTAMP
	);`)
	if err != nil {
		log.Fatalf("Failed to create urls table: %v", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS url_clicks (
		slug TEXT PRIMARY KEY,
		count INTEGER NOT NULL DEFAULT 0
	);`)
	if err != nil {
		log.Fatalf("Failed to create url_clicks table: %v", err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS click_analytics (
		id SERIAL PRIMARY KEY,
		slug TEXT,
		timestamp TIMESTAMP,
		ip TEXT,
		user_agent TEXT,
		referrer TEXT
	);`)
	if err != nil {
		log.Fatalf("Failed to create click_analytics table: %v", err)
	}

	// ------------------------------------------------------------------------------------
	// PERFORMANCE INDEXES
	// ------------------------------------------------------------------------------------

	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_urls_slug ON urls(slug);`)
	if err != nil {
		log.Fatalf("Failed to create idx_urls_slug: %v", err)
	}

	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_urlclicks_slug ON url_clicks(slug);`)
	if err != nil {
		log.Fatalf("Failed to create idx_urlclicks_slug: %v", err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_clickanalytics_slug ON click_analytics(slug);`)
	if err != nil {
		log.Fatalf("Failed to create idx_clickanalytics_slug: %v", err)
	}

	log.Println("Database initialized and tables + indexes ensured.")
}

func Save(slug string, longURL string, expiresAt *time.Time) error {
	_, err := db.Exec(`INSERT INTO urls (slug, long_url, expires_at) VALUES ($1, $2, $3)`, slug, longURL, expiresAt)
	if err != nil {
		log.Println("Error in Save: ", err)
	}
	return err
}


func Find(slug string) (string, *time.Time, error) {
	var longURL string
	var expiresAt *time.Time
	err := db.QueryRow(`SELECT long_url, expires_at FROM urls WHERE slug = $1`, slug).Scan(&longURL, &expiresAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil, errors.New("URL not found")
		}
		return "", nil, err
	}

	return longURL, expiresAt, nil
}

func Exists(slug string) (bool, error) {
	var exists bool
	err := db.QueryRow(`SELECT EXISTS(SELECT 1 FROM urls WHERE slug = $1)`, slug).Scan(&exists)
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
		return 0, expiresAt, err
	}
	return click, expiresAt, nil
}

// ------------------------------------------------------------------------------------

func DB() *sql.DB {
	return db
}
