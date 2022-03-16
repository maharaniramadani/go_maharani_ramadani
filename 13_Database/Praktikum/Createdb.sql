create database alta_online_shop;

-- create table soal no.2 a,b,c
MariaDB [alta_online_shop]> create table user (
    -> id integer not null auto_increment Primary key);
Query OK, 0 rows affected (0.075 sec)

MariaDB [alta_online_shop]> create table produk (
    -> id integer not null auto_increment Primary key);
Query OK, 0 rows affected (0.047 sec)

MariaDB [alta_online_shop]> create table produk_type (
    -> id integer not null auto_increment Primary key);
Query OK, 0 rows affected (0.041 sec)

MariaDB [alta_online_shop]> create table operator (
    -> id integer not null auto_increment Primary key);
Query OK, 0 rows affected (0.058 sec)

MariaDB [alta_online_shop]> create table produk_descriptions (
    -> id integer not null auto_increment Primary key);
Query OK, 0 rows affected (0.041 sec)

MariaDB [alta_online_shop]> create table payment_method (
    -> id integer not null auto_increment Primary key);

MariaDB [alta_online_shop]> create table transaksi (
    -> id integer not null auto_increment Primary key);
Query OK, 0 rows affected (0.047 sec)

MariaDB [alta_online_shop]> create table detail_transaksi (
    -> id integer not null auto_increment Primary key);
Query OK, 0 rows affected (0.043 sec)

-- soal no.3 creat table kurir
MariaDB [alta_online_shop]> create table kurir (
    -> id integer not null auto_increment Primary key,
    -> name varchar (25),
    -> created_at Timestamp Default current_Timestamp,
    -> updated_at Timestamp Default current_Timestamp);
Query OK, 0 rows affected (0.059 sec)

-- soal no.4 tambah ongkos_dasar pada column di table kurir
MariaDB [alta_online_shop]> alter table kurir ADD ongkos_dasar float;
Query OK, 0 rows affected (0.041 sec)
Records: 0  Duplicates: 0  Warnings: 0

-- soal no.5 rename table kurir
MariaDB [alta_online_shop]> alter table kurir RENAME TO shippings;
Query OK, 0 rows affected (0.021 sec)

-- soal no.6 hapus table shiping
MariaDB [alta_online_shop]> drop table shippings;
Query OK, 0 rows affected (0.017 sec)

-- soal no.7 relation
