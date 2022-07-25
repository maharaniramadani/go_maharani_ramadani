# 22 Middleware

## Resume
Dalam materi ini, mempelajari:


1. Echo Middleware
2. Type Echo Middleware
3. Implementation

### Pengertian

Middleware merupakan sebuah layer yang digunakan untuk memproses request-response HTTP sebelum diproses lebih lanjut dengan akses ke Echo#Context yang digunakan untuk melakukan tindakan tertentu, misalnya, mencatat setiap permintaan atau membatasi jumlah permintaan. contoh middleware yang disediakan oleh echo diantaranya adalah logger, basic auth, jwt.

Handler diproses pada akhirnya setelah semua middleware selesai dieksekusi.

Middleware yang terdaftar menggunakan Echo.Use() hanya dijalankan untuk jalur yang terdaftar setelah Echo.Use() dipanggil.


Echo.Pre() dapat digunakan untuk mendaftarkan middleware yang dieksekusi sebelum router memproses permintaan. Akan sangat membantu untuk membuat perubahan apa pun pada properti permintaan, misalnya, menambahkan atau menghapus garis miring dari jalur sehingga cocok dengan rute.