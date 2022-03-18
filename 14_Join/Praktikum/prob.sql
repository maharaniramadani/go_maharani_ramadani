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
    -> (3, 2, 1, 'LP1', 'Lemari Pakaian Portable', '20'),
    -> (4, 2, 1, 'LP2', 'Lemari Pakaian Plastik', '34'),
    -> (5, 2, 1, 'LP3', 'Lemari Rak Buku', '95');
Query OK, 3 rows affected (0.031 sec)
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
