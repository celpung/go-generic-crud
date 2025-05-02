# 🧩 go-generic-crud

`go-generic-crud` adalah package reusable untuk membuat endpoint CRUD generik menggunakan **Gin** dan **GORM** dengan prinsip Clean Architecture.

---

## ✨ Fitur

- Generic CRUD untuk entitas apapun (`Create`, `Read`, `ReadByID`, `Update`, `Delete`)
- Endpoint pencarian (`GET /search`) dan total data (`GET /count`)
- Middleware per metode HTTP
- Struktur Clean Architecture (Repository → Usecase → Delivery)
- Gampang digunakan dan diintegrasikan ke project Go mana pun

---

## 📦 Instalasi

```bash
go get github.com/celpung/go-generic-crud@latest
```

---

## 🗂️ Struktur Direkomendasikan untuk Project-mu

```
your-app/
├── entity/
│   └── slider.go
├── configs/
│   └── mysql/
│       └── connection.go
├── main.go
```

---

## 🧱 Contoh Entity (`entity/slider.go`)

```go
package entity

import "gorm.io/gorm"

type Slider struct {
    gorm.Model
    Title    string `json:"title"`
    ImageURL string `json:"image_url"`
}
```

---

## 🛠️ Contoh Setup Database GORM (`configs/mysql/connection.go`)

```go
package mysql

import (
    "your-app/entity"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    dsn := "user:password@tcp(localhost:3306)/yourdb?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Auto migrate your entities
    db.AutoMigrate(&entity.Slider{})

    DB = db
}
```

---

## 🌐 Setup Routing (`main.go`)

```go
package main

import (
    "reflect"

    "github.com/gin-gonic/gin"
    "github.com/celpung/go-generic-crud/crud_router"

    "your-app/entity"
    "your-app/configs/mysql"
)

func main() {
    r := gin.Default()

    mysql.Init()

    api := r.Group("/api")

    crud_router.SetupRouter[entity.Slider](
        api,
        mysql.DB,
        reflect.TypeOf(entity.Slider{}),
        "/sliders",
        map[string][]gin.HandlerFunc{
            "POST":   {AuthMiddleware()},
            "GET":    {},
            "PUT":    {AuthMiddleware()},
            "DELETE": {AuthMiddleware()},
        },
    )

    r.Run() // port 8080 by default
}
```

---

## 🔐 Contoh Middleware (`AuthMiddleware`)

```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != "Bearer mytoken" {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

---

## 📡 Daftar Endpoint Otomatis

| Method | Endpoint        | Deskripsi                   |
|--------|-----------------|-----------------------------|
| POST   | /sliders        | Tambah data                 |
| GET    | /sliders        | Ambil semua data            |
| GET    | /sliders/:id    | Ambil data berdasarkan ID   |
| PUT    | /sliders        | Update data                 |
| DELETE | /sliders/:id    | Hapus (soft delete)         |
| GET    | /sliders/search | Search berdasarkan keyword  |
| GET    | /sliders/count  | Hitung total data           |

---

## 📊 Query Parameters

- **Pagination, Sorting, dan Filtering**:

```
GET /sliders?page=1&limit=10&sortBy=created_at&title=Banner
```

- **Search global**:

```
GET /sliders/search?q=promo
```

---

## 🚀 Tips Penggunaan

- Buat folder `entity/` untuk semua struct entitas kamu.
- Gunakan middleware hanya jika dibutuhkan.
- Bisa digunakan untuk banyak entitas, misalnya:

```go
crud_router.SetupRouter[entity.User](api, mysql.DB, reflect.TypeOf(entity.User{}), "/users", nil)
crud_router.SetupRouter[entity.Book](api, mysql.DB, reflect.TypeOf(entity.Book{}), "/books", nil)
```

---

## 👷 Contribute

Pull request dipersilakan! Buat issue jika kamu ingin menambahkan fitur atau menemukan bug.

---

## 📝 Lisensi

MIT License © 2025