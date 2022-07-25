# 21 ORM and Code Structure (MVC)

## Resume
Dalam materi ini, mempelajari:


1. ORM
2. Database connection & migration
3. MVC

### Pengertian

object relational mapper (ORM) merupakan suatu teknik untuk mengubah data dari database dengan skema terkait menjadi object di suatu programming language dengan format yang sesuai
gorm merupakan library orm
kelebihan menggunakan orm: mengurangi perulangan statement query, langsung mengubah data menjadi siap pakai di suatu programmming language, cache query, screening data sebelum dikirim ke database
kekurangan menggunakan orm: menambah cost dari suatu proses karena harus melewati lapisan orm terlebih dahulu, memuat relasi yg tidak diperlukan, query yang kompleks bisa sangat ribet dituliskan menggunakan orm, bisa jadi ada sql function yang tidak didukung oleh library orm terkait
migration digunakan sebagai versioning dalam pengelolaan skema database
sebelum melakukan proses gorm, kita harus mendapatkan koneksi ke db terkait terlebih dahulu
kita perlu mendefinisikan skema suatu tabel ke dalam model
kita perlu mendefinisikan routes yang bisa diterima
kita perlu mendefinisikan handler untuk masing" routes