-- menghapus database yang lama
MariaDB [(none)]> drop database alta_online_shop;
Query OK, 1 row affected (0.002 sec)

-- membuat database baru
MariaDB [(none)]> create database alta_online_shop;
Query OK, 1 row affected (0.002 sec)

-- membuat table
CREATE TABLE `operators` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `created_at` timestamp,
);

CREATE TABLE `produk_type` (
  `id` int(11) PRIMARY KEY,
  `name` varchar(255),
  `created_at` timestamp
);

CREATE TABLE `produk` (
  `id` int PRIMARY KEY,
  `produk_type_id` int(11),
  `operator_id` int(11),
  `code` varchar(50),
  `name` varchar(100),
  `status` SMALLINT,
  `created_at` timestamp
);

CREATE TABLE `produk_descriptions` (
  `id` int(11) PRIMARY KEY,
  `produk_id` int(11),
  `description` text,
  `created_at` timestamp
);

CREATE TABLE `transaction_details` (
  `transaction_id` int(11),
  `produk_id` int(11),
  `status` varchar(10),
  `qyt` int(11),
  `price` numeric(25.2),
  `created_at` timestamp,
  PRIMARY KEY (`transaction_id`, `produk_id`)
);

CREATE TABLE `transaction` (
  `id` int(11) PRIMARY KEY,
  `user_id` int(11),
  `payment_method_id` int(11),
  `status` varchar(10),
  `total_qyt` int(11),
  `total_price` numeric(25,2),
  `created_at` timestamp
);

CREATE TABLE `payment_methods` (
  `id` int(11) PRIMARY KEY,
  `name` varchar(255),
  `status` SMALLINT,
  `created_at` timestamp
);

CREATE TABLE `users` (
  `id` int(11) PRIMARY KEY,
  `status` SMALLINT,
  `dob` date,
  `gender` char(1),
  `created_at` timestamp
);

ALTER TABLE `produk` ADD FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`);

ALTER TABLE `produk` ADD FOREIGN KEY (`produk_type_id`) REFERENCES `produk_type` (`id`);

ALTER TABLE `produk_descriptions` ADD FOREIGN KEY (`produk_id`) REFERENCES `produk` (`id`);

ALTER TABLE `transaction_details` ADD FOREIGN KEY (`produk_id`) REFERENCES `produk` (`id`);

ALTER TABLE `transaction_details` ADD FOREIGN KEY (`transaction_id`) REFERENCES `transaction` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transaction` ADD FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`);


-- 1.a
MariaDB [alta_online_shop]> insert into operators (id,name) values
    -> (1, 'pengguna 1'),
    -> (2, 'pengguna 2'),
    -> (3, 'pengguna 3'),
    -> (4, 'pengguna 4'),
    -> (5, 'pengguna 5');
Query OK, 5 rows affected (0.035 sec)
Records: 5  Duplicates: 0  Warnings: 0

-- 1.b
MariaDB [alta_online_shop]> insert into produk_type (id, name)
    -> values (1, 'skincare'),
    -> (2, 'lemari'),
    -> (3, 'baju');
Query OK, 3 rows affected (0.014 sec)
Records: 3  Duplicates: 0  Warnings: 0

-- 1.c
MariaDB [alta_online_shop]> insert into produk ( id, produk_type_id, operator_id, code, name, status) values
    -> (1, 1, 3, 'SC1', 'Toner', '200'),
    -> (2, 1, 3, 'SC2', 'Serum', '215');
Query OK, 2 rows affected (0.004 sec)
Records: 2  Duplicates: 0  Warnings: 0

-- 1.d
MariaDB [alta_online_shop]> insert into produk ( id, produk_type_id, operator_id, code, name, status) values
    -> (3, 2, 1, 'LP1', 'Lemari Pakaian Portable', '20'),
    -> (4, 2, 1, 'LP2', 'Lemari Pakaian Plastik', '34'),
    -> (5, 2, 1, 'LP3', 'Lemari Rak Buku', '95');
Query OK, 3 rows affected (0.031 sec)
Records: 3  Duplicates: 0  Warnings: 0

-- 1.e
MariaDB [alta_online_shop]> insert into produk ( id, produk_type_id, operator_id, code, name, status) values
    -> (6, 3, 4, 'BJ1', 'Baju Anak', '1000'),
    -> (7, 3, 4, 'BJ2', 'Dress', '140'),
    -> (8, 3, 4, 'BJ3', 'Kaos/T-shirt', '1450');
Query OK, 3 rows affected (0.008 sec)
Records: 3  Duplicates: 0  Warnings: 0

-- 1.f
MariaDB [alta_online_shop]> insert into produk_descriptions (id, produk_id, description) values
    -> (1, 1, 'Toner merupakan produk perawatan kulit untuk menghapus sisa makeup,kotoran,minyak dan mencerahkan wajah'),
    -> (2, 2, 'Serum merupakan produk perawatan kulit wajah untuk melembabkan, mencerahkan, hingga mengatasi kerutan'),
    -> (3, 3, 'Lemari pakaian portable berbahan kain yang bisa dilipat'),
    -> (4, 4, 'Lemari pakaian plastik berkualitas tinggi ukuran 2x3 tingkat'),
    -> (5, 5, 'Lemari Buku, lemari portable serbaguna untuk merapikan buku, tempat majalah atau dokumen'),
    -> (6, 6, 'Baju anak usia 1-5 tahun bahan katun'),
    -> (7, 7, 'Dress cantik berkualitas premium'),
    -> (8, 8, 'Kaos/Tshirt oblong berkualitas premium overzise');
Query OK, 8 rows affected (0.030 sec)
Records: 8  Duplicates: 0  Warnings: 0

-- 1.g
MariaDB [alta_online_shop]> insert into payment_methods (id, name, status) values
    -> (1, 'altapay', 200),
    -> (2, 'E-Banking', 200),
    -> (3, 'Transfer ATM', 200);
Query OK, 3 rows affected (0.009 sec)
Records: 3  Duplicates: 0  Warnings: 0

-- 1.h
MariaDB [alta_online_shop]> insert into users (id, status, dob, gender) values
    -> (1, 200, '2001-11-17', 'F'),
    -> (2, 200, '2001-03-29', 'F'),
    -> (3, 200, '2001-02-02', 'M'),
    -> (4, 200, '2001-12-09', 'M'),
    -> (5, 200, '2000-10-20', 'F');
Query OK, 5 rows affected (0.004 sec)
Records: 5  Duplicates: 0  Warnings: 0

-- 1.i
MariaDB [alta_online_shop]> insert into transaction (id, user_id, payment_method_id, status, total_qyt, total_price) values
    -> (1,1,1,'sukses',2,90000),
    -> (2,1,2,'sukses',4,220000),
    -> (3,1,3,'sukses',3,560000),
    -> (4,2,1,'tunda',5,755000),
    -> (5,2,2,'sukses',1,45000),
    -> (6,2,3,'batal',3,260000),
    -> (7,3,1,'batal',6,800000),
    -> (8,3,2,'sukses',1,80000),
    -> (9,3,3,'sukses',2,100000),
    -> (10,4,1,'tunda',3,110000),
    -> (11,4,2,'sukses',5,225000),
    -> (12,4,3,'tunda',2,110000),
    -> (13,5,1,'sukses',4,100000),
    -> (14,5,2,'sukses',5,755000),
    -> (15,5,3,'tunda',4,495000);
Query OK, 15 rows affected (0.010 sec)
Records: 15  Duplicates: 0  Warnings: 0

-- 1.j
MariaDB [alta_online_shop]> insert into transaction_details (transaction_id, produk_id, status, qyt, price)
    -> values (1, 1, 'selesai', 2, 90000),
    -> (1, 2, 'dikirim', 1, 45000),
    -> (1, 3, 'dikirim', 4, 220000),
    -> (2, 4, 'selesai', 3, 200000),
    -> (2, 5, 'selesai', 1, 110000),
    -> (2, 6, 'selesai', 4, 590000),
    -> (3, 7, 'selesai', 3, 900000),
    -> (3, 8, 'selesai', 5, 900000),
    -> (3, 1, 'selesai', 1, 100000),
    -> (4, 2, 'selesai', 3, 500000),
    -> (4, 3, 'selesai', 2, 80000),
    -> (4, 4, 'dikirim', 2, 900000),
    -> (5, 5, 'selesai', 2, 300000),
    -> (5, 6, 'selesai', 5, 350000),
    -> (5, 7, 'selesai', 3, 650000),
    -> (6, 8, 'selesai', 2, 300000),
    -> (6, 1, 'selesai', 2, 110000),
    -> (6, 2, 'selesai', 2, 120000),
    -> (7, 3, 'dikirim', 4, 900000),
    -> (7, 4, 'dikirim', 5, 1500000),
    -> (7, 5, 'selesai', 1, 450000),
    -> (8, 6, 'selesai', 12, 540000),
    -> (8, 7, 'selesai', 12, 540000),
    -> (8, 8, 'selesai', 24, 2400000),
    -> (9, 1, 'selesai', 4, 220000),
    -> (9, 2, 'selesai', 4, 220000),
    -> (9, 3, 'selesai', 2, 900000),
    -> (10, 4, 'dikirim', 2, 1000000),
    -> (10, 5, 'selesai', 1, 450000),
    -> (10, 6, 'selesai', 1, 500000),
    -> (11, 7, 'selesai', 1, 150000),
    -> (11, 8, 'selesai', 5, 450000),
    -> (11, 1, 'selesai', 3, 450000),
    -> (12, 2, 'selesai', 5, 325000),
    -> (12, 3, 'selesai', 1, 350000),
    -> (12, 4, 'dikirim', 5, 2250000),
    -> (13, 5, 'dikirim', 8, 2400000),
    -> (13, 6, 'selesai', 12, 540000),
    -> (13, 7, 'selesai', 12, 700000),
    -> (13, 8, 'selesai', 5, 450000),
    -> (14, 1, 'selesai', 4, 180000),
    -> (14, 2, 'selesai', 8, 520000),
    -> (14, 3, 'dikirim', 8, 2400000),
    -> (15, 4, 'selesai', 1, 450000),
    -> (15, 5, 'selesai', 1, 500000),
    -> (15, 6, 'selesai', 24, 2500000);
Query OK, 46 rows affected (0.027 sec)
Records: 46  Duplicates: 0  Warnings: 0
