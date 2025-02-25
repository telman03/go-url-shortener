# URL Shortener&#x20;

## 📌 Overview

This is a high-performance \*\*URL shortener \*\* built with **Golang (Fiber)**, **PostgreSQL**, **Redis**, and **Docker**. It allows users to shorten long URLs, track click analytics, and retrieve stored URLs efficiently.

## 🚀 Features

- **Shorten URLs**: Generate a unique short code for any long URL.
- **Redirect to Original URL**: Use the short code to access the original link.
- **Caching with Redis**: Improve performance by caching URLs.
- **Analytics Tracking**: Monitor how many times a short URL has been accessed.



## 🏗 Tech Stack

- **Golang (Fiber)** - Fast and lightweight web framework
- **PostgreSQL** - Database for storing URLs
- **Redis** - Caching layer for fast lookups
- **Docker & Docker Compose** - Containerized deployment

## 📦 Installation & Setup

### 1️⃣ Clone the Repository

```sh
git clone https://github.com/telman03/go-url-service.git
cd go-url-service
```

### 2️⃣ Set Up Environment Variables

Create a `.env` file and configure the database & Redis:

```env
DATABASE_URL=postgres://user:password@localhost:5432/urlshortener
REDIS_ADDR=localhost:6379
```

### 3️⃣ Run Redis & PostgreSQL using Docker

```sh
docker-compose up -d redis postgres
```

### 4️⃣ Run the Service

```sh
go run main.go
```

## 📡 API Endpoints

### 1️⃣ Shorten a URL

```http
POST /api/shorten
```

**Request Body:**

```json
{
  "url": "https://example.com"
}
```

**Response:**

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### 2️⃣ Redirect to Original URL

```http
GET /:shortcode
```

Example: `GET http://localhost:8080/abc123` → Redirects to `https://example.com`

### 3️⃣ Get URL Analytics

```http
GET /api/stats/:shortcode
```

**Response:**

```json
{
  "shortcode": "abc123",
  "clicks": 10
}
```

## 🛠️ Future Improvements

- 🔑 **User authentication** to manage links
- 📅 **URL expiration feature**
- 📊 **Admin dashboard** for analytics
- 🌍 **Custom domains support**

---

💡 **Contributions are welcome!** Feel free to open an issue or submit a pull request. 🚀

