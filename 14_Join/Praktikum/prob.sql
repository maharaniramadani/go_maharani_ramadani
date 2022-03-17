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