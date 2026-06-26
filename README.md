# 🚗 SpotSync API

A Smart Parking & EV Charging Reservation System built with Go (Golang), Echo, GORM, and PostgreSQL.

## 📌 Project Status

🚧 Currently under development.

## 🛠️ Tech Stack

* Go (Golang)
* Echo Framework
* GORM
* PostgreSQL
* JWT Authentication
* bcrypt
* Validator

## 📁 Project Structure

```text
.
├── cmd/
├── config/
├── dto/
├── handler/
├── middleware/
├── models/
├── repository/
├── routes/
├── service/
├── utils/
├── .env
├── go.mod
└── README.md
```

## 🚀 Features (Planned)

* User Registration
* User Login (JWT Authentication)
* Parking Zone Management
* EV Charging Reservation
* Parking Reservation
* Cancel Reservation
* Role-based Authorization (Admin & Driver)
* Prevent Overbooking using Database Transactions
* RESTful API
* Clean Architecture

## ▶️ Run Locally

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/Web-Dev-B6A6.git
```

Go to the project folder:

```bash
cd Web-Dev-B6A6
```

Install dependencies:

```bash
go mod tidy
```

Run the server:

```bash
go run ./cmd/server
```

The server will start at:

```text
http://localhost:8080
```

## 📅 Development Progress

* [x] Project Setup
* [x] Echo Server
* [ ] Database Configuration
* [ ] User Authentication
* [ ] JWT Middleware
* [ ] Parking Zone Module
* [ ] Reservation Module
* [ ] Concurrency Handling
* [ ] Deployment

## 👨‍💻 Author

Al Amin
