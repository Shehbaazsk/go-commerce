
# 🛒 Go-Commerce API

A scalable, modular eCommerce RESTful API built with **Go (Gin framework)**.  
Implements **JWT authentication**, **role-based access control (RBAC)** at the route level, and follows a clean, app-based, layered architecture with separated concerns.

---

## 📦 Features

✅ Modular folder structure (MVC + service layer)  
✅ JWT-based authentication  
✅ Custom role-based access control (admin, customer, etc.)  
✅ Public and protected API routes  
✅ Custom middlewares for authentication and RBAC  
✅ Environment-based configuration  
✅ Docker + Docker Compose for containerization  
✅ Password hashing with bcrypt  
✅ Product, Order, and User management APIs  
✅ Ready for Swagger integration  

---

## 📂 Project Structure

```bash
go-commerce/
├── cmd/                    # Application entrypoint (main.go)
├── config/                 # Database & environment configuration
├── controllers/            # HTTP handler functions
├── middlewares/            # JWT auth and RBAC middleware
├── models/                 # Data models (GORM structs)
├── routes/                 # Route grouping and middleware registration
├── services/               # Business logic and database interaction
├── tests/                  # Unit/integration tests
├── utils/                  # Helper utilities (JWT, password hashing)
├── go.mod                  # Go module definition
├── .env                    # Environment variables
└── README.md               # Project documentation
```

---

## 🛠️ Tech Stack

- [Go 1.22+](https://go.dev)
- [Gin Web Framework](https://gin-gonic.com)
- [GORM](https://gorm.io) (ORM)
- [PostgreSQL](https://www.postgresql.org)
- JWT for authentication
- bcrypt for password hashing

---

## 🚀 Getting Started

### Prerequisites

- Go 1.22+

---

### 📥 Clone the repository

```bash
git clone https://github.com/shehbaazsk/go-commerce.git
cd go-commerce
```

---

### ⚙️ Configure environment variables

Create a `.env` file in the project root:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=go_commerce
JWT_SECRET=your_jwt_secret_key
```

---


## 📖 API Endpoints

| Method | Endpoint               | Access  | Description               |
|:--------|:-----------------------|:-----------|:---------------------------|
| POST   | `/api/register`        | Public   | User registration         |
| POST   | `/api/login`           | Public   | User login                 |
| GET    | `/api/products`        | Public   | List products              |
| POST   | `/api/products`        | Admin    | Create product             |
| GET    | `/api/profile`         | Customer | View own profile           |
| GET    | `/api/admin/users`     | Admin    | List all users             |

👉 Full route map is in `/routes/routes.go`

---

## 📚 License

MIT License.  
Made with ❤️ in Go.

---

## 📬 Contact

- GitHub: [shehbaaz shaikh](https://github.com/shehbaazsk)
- Email: shehbaazwebdev@gmail.com
