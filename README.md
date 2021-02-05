# CRUD

Project ini merupakan project yang dibuat untuk latihan saja

## Mulai

Berikut ini project sederhana untuk melakukan CRUD Product dengan beberapa relasi tabel. Bahasa yang digunakan adalah Go.

### Persiapan

Beberapa yang perlu dilakukan sebelum menjalankan service ini

```
1. Buat USER pada postgresql dengan nama 'postgres' password 'postgres'
2. Buat .env file seperti contoh .env.example
2. Import Database sql yang sudah disediakan
3. Jika ingin merubah config db cukup dirubah di file env yang telah disediakan
4. Lanjut ke step berikutnya
```

### Menjalankan service

1. Silahkan mengikuti intruksi menginstall Go kedalam OS anda pada laman berikut [ini] (https://golang.org/doc/install)

2. Masuk ke direktori project lalu jalankan binary file dengan cara 
```
go run dockerize-go-app
```

### Cara lain
Jika anda ingin mencoba menjalankan projectnya tanpa binary file lakukan hal birkut ini

1. Masuk ke folder project jalankan perintah berikut ini
```
go mod
```

2. Setelah selesai mengunduh dependency yang diperlukan selanjutnya compile project dengan perintah berikut
```
go build
```

3. File binary dari nama project sudah terbentuk, selanjutnya jalankan file binary tersebut
```
go run nama_binary_file
```

**NB** : nama binary file sesuai dengan nama project

Jika ingin menggunakan `Dockerfile` bisa gunakan file yang telah disediakan. `Dockerfile` yang digunakan di proejct ini menggunakan pattern multistage build.

## Mencoba service

Untuk mencoba service silahkan import file json di dalam folder postman-collection ke aplikasi postman

## Dibangun dengan

* [Echo Framework](https://echo.labstack.com/) - Web framework yang digunakan

## Penulis

* **Manggala Pramuditya Wiryawan** - *Inisial* - [Wiryawan46](https://github.com/wiryawan46)
