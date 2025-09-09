# Ethio-Football API ⚽

Backend API for Ethiopian Premier League (ETH) and English Premier League (EPL) — providing fixtures, standings, and team data for fans and developers.

Built with **Go (Gin)** and **Redis** for caching.

---

## ✨ Features
- 📅 Fixtures with logos, status, and scores  
- 📊 Standings by league and season  
- 🏟 Team bios and curated data (Ethiopian clubs)  
- 🔐 (Optional) Authentication endpoints for future versions  
- ⚡ Caching with Redis (faster responses, offline-friendly)

---

## 📂 Project Structure
```

.
├── Delivery/       # Controller and router
  ├── main.go         # Entry point
├── Domain/         # Core models (Team, Fixture, Standing, etc.)
├── Repository/     # Repositories (Redis + API fetchers)
├── Infrastructure/ # API integration (API-Football, others)
├── Usecase/        # Business logic (fixtures, teams)

````

---

## 🚀 Getting Started

### Prerequisites
- [Go](https://go.dev/) 1.21+
- [Redis](https://redis.io/)
- API key from [API-Football](https://www.api-football.com/)

### Setup
```bash
# Clone the repo
git clone https://github.com/surafelbkassa/ethiofootball-api.git
cd ethiofootball-api

# Install dependencies
go mod tidy
````

### Environment Variables (`.env`)

```env
REDIS_ADDRESS=""
REDIS_USERNAME=""
REDIS_PASSWORD=""
API_FOOTBALL_KEY=""
```

### Run the server

```bash
go run main.go
```

---

## 🔑 API Endpoints

### Fixtures

```
GET /fixtures?league=EPL&season=2023&from=2024-05-01&to=2024-05-10
```

### Standings

```
GET /standings?league=EPL&season=2023
```

### Teams

```
GET /team/:id/bio
```

*(See full docs in Postman collection)*

---

## 🛠 Tech Stack

* **Go (Gin)** — HTTP server
* **Redis** — caching & persistence
* **API-Football** — fixtures & standings data

---

## 👥 Contributors

* Backend Team (A2SV) — [EthioFootball Project](https://github.com/abrshodin/ethio-fb-backend)
* Fork maintained by **@surafelbkassa**

---

## 📜 License

MIT — feel free to use, extend, or contribute.

