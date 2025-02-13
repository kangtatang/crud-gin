# Gin User Management

Gin User Management adalah aplikasi CRUD berbasis Golang menggunakan Gin dan GORM dengan PostgreSQL sebagai database.

Aplikasi **Gin User Management** adalah REST API sederhana yang dibangun menggunakan **Golang**, **Gin**, dan **GORM** untuk mengelola pengguna. API ini mencakup operasi CRUD (Create, Read, Update, Delete) pada entitas pengguna, saat ini tanpa menggunakan middleware.

Aplikasi ini dibuat sebagai bahan referensi belajar bagi pemula yang sedang belajar Golang dengan menggunakan Gin Framework. 

## 🚀 Fitur
- **Membuat pengguna baru** (Create User)
- **Mendapatkan daftar pengguna** (Get Users)
- **Mendapatkan detail pengguna berdasarkan ID** (Get User by ID)
- **Memperbarui data pengguna** (Update User)
- **Menghapus pengguna** (Delete User)
- **Konfigurasi dengan .env: Menggunakan environment variables**.
- **Auto Migrate Database: Migrasi otomatis saat aplikasi dijalankan**.

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

## ⚙️ Konfigurasi
Buat file `.env` di root project dengan isi berikut:

```ini
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=yourdbname
DB_PORT=5432
APP_PORT=8080
```

## 🛠 Cara Menjalankan di Windows
```sh
go run main.go
```


## 🚀 Build dan Run Aplikasi

### **1️⃣ Build untuk Windows**
```sh
go build -o gin-user-app.exe main.go
```
Jalankan:
```sh
gin-user-app.exe
```

### **2️⃣ Build untuk Linux (Cross-Compilation di Windows)**
```sh
set GOOS=linux
set GOARCH=amd64
go build -o gin-user-app main.go
```
Transfer file ke server Linux:
```sh
scp gin-user-app user@your-server-ip:/home/user/
```

### **3️⃣ Jalankan Aplikasi di Linux**
```sh
ssh user@your-server-ip
cd /home/user/
chmod +x gin-user-app
./gin-user-app
```

### **4️⃣ Jalankan Aplikasi sebagai Service (Opsional)**
Buat file systemd:
```sh
sudo nano /etc/systemd/system/gin-user.service
```
Isi dengan:
```ini
[Unit]
Description=Gin User Management Service
After=network.target

[Service]
ExecStart=/home/user/gin-user-app
WorkingDirectory=/home/user
Restart=always
User=user

[Install]
WantedBy=multi-user.target
```
Simpan dan aktifkan:
```sh
sudo systemctl daemon-reload
sudo systemctl enable gin-user.service
sudo systemctl start gin-user.service
```

## 📜 Lisensi
Proyek ini berlisensi di bawah **MIT License**.

## 🤝 Kontribusi
Silakan buat **Pull Request** jika ingin berkontribusi atau melaporkan masalah.

>kontak saya di : kang.tatang@yahoo.co.id
