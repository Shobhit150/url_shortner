# 🔗 URL Shortener (Scalable & Zero-Tolerance)

A **highly scalable**, **zero-tolerance** URL shortener built with **Go**, **PostgreSQL**, and **Redis**. Designed to be production-ready with support for fault tolerance, high availability, and extensibility.

---

## 🚀 Features

- Convert long URLs to short, unique slugs
- Fast redirection using Redis cache
- Unique slug generation with collision handling
- No data loss (strong DB consistency)
- Rate-limiting and retries
- Ready for containerization and cloud deployment
- Extendable for analytics, expiry, QR codes, etc.

---

## 🧱 Project Structure

url-shortener/
├── cmd/ # Entry point for the HTTP/GRPC server
│ └── server/
│ └── main.go
├── internal/
│ ├── handler/ # REST/GRPC route handlers
│ ├── service/ # Business logic layer
│ ├── repository/ # Database interaction
│ ├── cache/ # Redis cache interface
│ └── utils/ # Slug generator, validators, etc.
├── proto/ # GRPC protobufs (optional)
├── migrations/ # SQL migration files
├── configs/ # Config and environment files
├── Dockerfile # Docker setup for the service
├── docker-compose.yml # Dev orchestration with Redis and Postgres
├── go.mod
└── README.md


---

## ⚙️ Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Cache:** Redis
- **Containerization:** Docker
- **API:** REST (option to switch to gRPC)
- **Monitoring:** Prometheus + Grafana (optional)
- **Queue (optional):** Kafka/NATS for analytics

---