# Go JWT Authentication System

Sistem Autentikasi JWT sederhana yang dibangun menggunakan Go, mengikuti standar **Standard Go Project Layout**. Project ini mencakup fitur registrasi, login, manajemen database dengan migrasi, dan validasi data.

## 🚀 Fitur Utama

- **User Registration**: Membuat akun baru dengan password yang di-hash menggunakan `bcrypt`.
- **User Login**: Autentikasi user (JWT holder).
- **Database Migration**: Manajemen skema database MySQL menggunakan `golang-migrate`.
- **Validation**: Validasi input payload menggunakan `go-playground/validator`.
- **Environment Support**: Konfigurasi menggunakan file `.env`.
- **Unit Testing**: Pengujian handler menggunakan `testify` dan `httptest`.

## 📁 Struktur Folder

```text
├── bin/                # Binary hasil build
├── cmd/
│   ├── api/
│   │   └── main.go     # Entry point aplikasi API
│   └── migrate/
│       ├── main.go     # Runner untuk migrasi database
│       └── migrations/ # File SQL migrasi (.up.sql dan .down.sql)
├── internal/
│   ├── api/            # Logika server dan routing utama
│   ├── config/         # Manajemen konfigurasi (.env)
│   ├── db/             # Inisialisasi koneksi database
│   ├── service/        # Business logic per fitur (User)
│   ├── types/          # Definisi struct dan interface
│   └── utils/          # Fungsi pembantu (JSON, Hash, dll)
├── .env                # File konfigurasi (DB Credentials)
├── Makefile            # Automasi task (run, build, migrate)
└── go.mod
```

## 🛠️ Cara Menjalankan

### 1. Prasyarat

- Go 1.25+
- MySQL Server sudah berjalan

### 2. Setup Environment

Buat file `.env` di root folder:

```env
DB_USER=root
DB_PASSWD=password_anda
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_starter_project
```

### 3. Jalankan Migrasi

Pastikan database `go_starter_project` sudah ada di MySQL, lalu jalankan:

```bash
make migrate-up
```

### 4. Jalankan Aplikasi

```bash
make run
```

Server akan berjalan di `http://localhost:8080`.

## 🧪 Testing

Jalankan semua unit test dengan perintah:

```bash
make test
```

## 🔌 API Endpoints

| Method | Endpoint           | Keterangan             |
| :----- | :----------------- | :--------------------- |
| `POST` | `/api/v1/register` | Mendaftarkan user baru |
| `POST` | `/api/v1/login`    | Login user             |

---

**Author**: Fikri Lazuardi
