# Gin User Management

Aplikasi **Gin User Management** adalah REST API sederhana yang dibangun menggunakan **Golang**, **Gin**, dan **GORM** untuk mengelola pengguna. API ini mencakup operasi CRUD (Create, Read, Update, Delete) pada entitas pengguna, saat ini tanpa menggunakan middleware.

Aplikasi ini dibuat sebagai bahan referensi belajar bagi pemula yang sedang belajar Golang dengan menggunakan Gin Framework. 

## 🚀 Fitur
- **Membuat pengguna baru** (Create User)
- **Mendapatkan daftar pengguna** (Get Users)
- **Mendapatkan detail pengguna berdasarkan ID** (Get User by ID)
- **Memperbarui data pengguna** (Update User)
- **Menghapus pengguna** (Delete User)

## 🛠️ Teknologi yang Digunakan
- **Golang** sebagai bahasa pemrograman utama
- **Gin** sebagai web framework
- **GORM** sebagai ORM untuk database
- **PostgreSQL** sebagai database utama

## 📦 Instalasi
1. **Clone repository ini**
   ```sh
   git clone https://github.com/kangtatang/gin-user-management.git
   cd gin-user-management
   ```
2. **Instal dependency**
   ```sh
   go mod tidy
   ```
3. **Konfigurasi database**
   Buat file `.env` dan sesuaikan dengan konfigurasi database PostgreSQL:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=password
   DB_NAME=gin_users
   ```

4. **Jalankan aplikasi**
   ```sh
   go run main.go
   ```

## 📌 Endpoint API
| Method | Endpoint        | Deskripsi               |
|--------|----------------|-------------------------|
| **GET**    | `/users`        | Mendapatkan daftar pengguna |
| **GET**    | `/users/:id`    | Mendapatkan detail pengguna |
| **POST**   | `/users`        | Menambahkan pengguna baru |
| **PUT**    | `/users/:id`    | Memperbarui pengguna |
| **DELETE** | `/users/:id`    | Menghapus pengguna |

## 📝 Contoh Payload
### **1. Create User**
**Request:**
```json
{
    "name": "John Doe",
    "email": "john.doe@example.com"
}
```
**Response:**
```json
{
    "status": 201,
    "message": "User created successfully",
    "data": {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "created_at": 1700000000,
        "updated_at": 1700000000
    }
}
```

### **2. Get Users**
**Response:**
```json
{
    "status": 200,
    "message": "Users retrieved successfully",
    "data": [
        {
            "id": 1,
            "name": "John Doe",
            "email": "john.doe@example.com",
            "created_at": 1700000000,
            "updated_at": 1700000000
        }
    ]
}
```

## 🏗️ Struktur Proyek
```
├── database/
│   ├── database.go   # Inisialisasi database PostgreSQL
│
├── handlers/
│   ├── user_handler.go   # Handler untuk operasi CRUD pengguna
│
├── helpers/
│   ├── response.go   # Helper untuk standar respons API
│
├── models/
│   ├── user.go   # Struktur model User
│
├── main.go   # Entry point aplikasi
├── go.mod   # Dependency management
├── README.md   # Dokumentasi proyek
```

## 📜 Lisensi
Proyek ini berlisensi di bawah **MIT License**.

## 🤝 Kontribusi
Silakan buat **Pull Request** jika ingin berkontribusi atau melaporkan masalah.

>kontak saya di : kang.tatang@yahoo.co.id
