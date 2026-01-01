# CLI Task Manager

CLI Task Manager adalah aplikasi command-line sederhana yang dibangun menggunakan **Go (Golang)** untuk mengelola task (create, update, delete, list) dengan penyimpanan data berbasis file JSON.

Project ini dibuat sebagai media belajar Go dengan pendekatan berbasis project dan struktur kode yang rapi serta mudah dikembangkan.

---

## Fitur Saat Ini

- Create, update, delete, dan list task melalui CLI
- Penyimpanan data menggunakan file JSON
- Generate ID otomatis

## Fitur yang Sedang Dikembangkan

- Parsing tanggal dengan format `YYYY-MM-DD`
- Menampilkan tanggal task dibuat
- Reset data untuk memulai dari awal

---

## Prasyarat

- Go versi 1.20 atau lebih baru

---

## Cara Menjalankan (Run)

Jalankan langsung menggunakan Go:

```bash
go run ./cmd/tasker <command> [argument]
```
Add Task:

```bash
go run ./cmd/tasker add "judul task"
```

List Tasks:
```bash
go run ./cmd/tasker list
```

Update Task:
```bash
go run ./cmd/tasker update <id>
```

Delete Task:
```bash
go run ./cmd/tasker delete <id>
```

Markdone Task:
```bash
go run ./cmd/tasker done <id> (for markdone task)
```

## Cara Menjalankan (Build)

Jalankan build file:

```bash
go build -o tasker
```

Jika sudah di compile:
```bash
./tasker <command> [argument]
