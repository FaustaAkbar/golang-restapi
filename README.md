# Golang REST API

REST API sederhana berbasis [Go](https://go.dev/) untuk mengelola data user dengan [Gin](https://gin-gonic.com/), [GORM](https://gorm.io/), MySQL, dan Redis. Proyek ini menggunakan arsitektur berlapis untuk memisahkan routing, handler, service, repository, dan model.

## Fitur

- REST API dengan framework Gin
- Penyimpanan data user menggunakan MySQL dan GORM
- Cache daftar user menggunakan Redis selama 5 menit
- Migrasi tabel otomatis saat aplikasi dijalankan
- Konfigurasi melalui environment variable
- Struktur proyek berlapis
- Infrastruktur lokal MySQL dan Redis melalui Docker Compose

## Teknologi

- Go 1.26
- Gin
- GORM
- MySQL 8
- Redis 7
- Docker Compose
- godotenv

## Struktur Proyek

```text
golang_restapi/
├── cmd/
│   └── main.go                 # Entry point aplikasi
├── internal/
│   ├── config/                 # Pembacaan environment variable
│   ├── database/               # Koneksi MySQL dan Redis
│   ├── handler/                # HTTP handler
│   ├── models/                 # Model/database schema
│   ├── repository/             # Akses database
│   ├── routes/                 # Registrasi endpoint
│   └── service/                # Business logic dan cache
├── docker-compose.yml          # MySQL dan Redis untuk development
├── go.mod
└── go.sum
```

## Prasyarat

- [Go](https://go.dev/dl/) 1.26 atau lebih baru
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) atau Docker Engine
- Git

## Instalasi dan Menjalankan

### 1. Clone repository

```bash
git clone https://github.com/USERNAME/golang_restapi.git
cd golang_restapi
```

Ganti `USERNAME` dengan username GitHub Anda.

### 2. Buat file `.env`

Buat file `.env` dari contoh berikut:

```dotenv
MYSQL_DSN=root:root@tcp(localhost:3306)/app_db?charset=utf8mb4&parseTime=True&loc=Local
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
```

Untuk development lokal, nilai di atas dapat digunakan bersama `docker-compose.yml`.

> Jangan commit file `.env` ke repository. File ini sudah dicakup oleh `.gitignore`.

### 3. Jalankan MySQL dan Redis

```bash
docker compose up -d
```

Perintah tersebut menjalankan:

| Service | Port |
| --- | --- |
| MySQL | `3306` |
| Redis | `6379` |

### 4. Jalankan aplikasi

```bash
go mod download
go run ./cmd
```

Server berjalan di:

```text
http://localhost:8080
```

Saat startup, aplikasi akan membuat tabel `users` secara otomatis melalui GORM AutoMigrate.

### Menjalankan sebagai binary

```bash
go build -o golang-restapi.exe ./cmd
./golang-restapi.exe
```

## Konfigurasi

| Variable | Deskripsi | Contoh |
| --- | --- | --- |
| `MYSQL_DSN` | Connection string MySQL untuk GORM | `root:root@tcp(localhost:3306)/app_db?charset=utf8mb4&parseTime=True&loc=Local` |
| `REDIS_ADDR` | Alamat Redis | `localhost:6379` |
| `REDIS_PASSWORD` | Password Redis; kosongkan bila tidak digunakan | _(kosong)_ |

## Daftar Endpoint

Base URL: `http://localhost:8080/api/v1`

| Method | Endpoint | Deskripsi |
| --- | --- | --- |
| `GET` | `/users` | Mengambil semua user |
| `POST` | `/users` | Membuat user baru |

### Mendapatkan semua user

```bash
curl http://localhost:8080/api/v1/users
```

Contoh respons:

```json
[
  {
    "ID": 1,
    "name": "Fauzan",
    "Email": "fauzan@example.com"
  }
]
```

Respons pertama dapat berasal dari MySQL. Respons berikutnya berasal dari cache Redis selama 5 menit.

### Membuat user

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Fauzan\",\"email\":\"fauzan@example.com\"}"
```

Contoh respons `201 Created`:

```json
{
  "ID": 1,
  "name": "Fauzan",
  "Email": "fauzan@example.com"
}
```

Email bersifat unique pada database. Saat user berhasil dibuat, cache key `users` akan dihapus agar data berikutnya diambil ulang dari MySQL.

## Alur Redis Cache

1. `GET /api/v1/users` memeriksa cache key `users`.
2. Jika cache tersedia, data dikembalikan langsung dari Redis.
3. Jika cache tidak tersedia, data diambil dari MySQL.
4. Hasil MySQL disimpan ke Redis selama 5 menit.
5. `POST /api/v1/users` menghapus cache `users` setelah data berhasil disimpan.

## Docker

Menjalankan semua dependency:

```bash
docker compose up -d
```

Melihat status container:

```bash
docker compose ps
```

Melihat log:

```bash
docker compose logs -f
```

Menghentikan dan menghapus container:

```bash
docker compose down
```

## Pengembangan

Memformat source code:

```bash
go fmt ./...
```

Memeriksa dependency dan kompilasi:

```bash
go mod tidy
go test ./...
go build ./cmd
```

## Catatan

- Server saat ini berjalan di port `8080`.
- Redis tidak memiliki password pada konfigurasi Docker Compose bawaan.
- Password MySQL development pada contoh adalah `root`; gunakan kredensial yang aman untuk production.
- Proyek belum memiliki endpoint `GET /users/:id`, `PUT`, dan `DELETE`.
- Proyek belum memiliki authentication, authorization, atau validation schema yang lengkap.

## Kontribusi

1. Fork repository ini.
2. Buat branch fitur: `git checkout -b feature/nama-fitur`.
3. Commit perubahan: `git commit -m "feat: tambah fitur baru"`.
4. Push branch: `git push origin feature/nama-fitur`.
5. Buat Pull Request.

## Lisensi

Tambahkan lisensi proyek Anda di sini, misalnya MIT, Apache 2.0, atau GPL-3.0.
