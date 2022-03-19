-- 1
MariaDB [alta_online_shop]> select * from transaction where user_id = 1 UNION select * from transaction where user_id = 2 ;
+----+---------+-------------------+--------+-----------+-------------+---------------------+
| id | user_id | payment_method_id | status | total_qyt | total_price | created_at          |
+----+---------+-------------------+--------+-----------+-------------+---------------------+
|  1 |       1 |                 1 | sukses |         2 |    90000.00 | 2022-03-18 00:25:44 |
|  2 |       1 |                 2 | sukses |         4 |   220000.00 | 2022-03-18 00:25:44 |
|  3 |       1 |                 3 | sukses |         3 |   560000.00 | 2022-03-18 00:25:44 |
|  4 |       2 |                 1 | tunda  |         5 |   755000.00 | 2022-03-18 00:25:44 |
|  5 |       2 |                 2 | sukses |         1 |    45000.00 | 2022-03-18 00:25:44 |
|  6 |       2 |                 3 | batal  |         3 |   260000.00 | 2022-03-18 00:25:44 |
+----+---------+-------------------+--------+-----------+-------------+---------------------+
6 rows in set (0.036 sec)

-- 2

MariaDB [alta_online_shop]> select SUM(total_price) from transaction where user_id = 1;
+------------------+
| SUM(total_price) |
+------------------+
|        870000.00 |
+------------------+
1 row in set (0.005 sec

-- 3
MariaDB [alta_online_shop]> SELECT COUNT(transaction_id) FROM transaction_details WHERE produk_id IN (SELECT id FROM produk WHERE produk_type_id = 2);
+-----------------------+
| COUNT(transaction_id) |
+-----------------------+
|                    18 |
+-----------------------+
1 row in set (0.035 sec)

-- 4

MariaDB [alta_online_shop]> select p.*,pt.name AS produk_type_name  from produk p INNER JOIN produk_type pt ON p.produk_type_id = pt.id;
+----+----------------+-------------+------+-------------------------+--------+---------------------+------------------+
| id | produk_type_id | operator_id | code | name                    | status | created_at          | produk_type_name |
+----+----------------+-------------+------+-------------------------+--------+---------------------+------------------+
|  3 |              2 |           1 | LP1  | Lemari Pakaian Portable |     20 | 2022-03-17 23:08:02 | lemari           |
|  4 |              2 |           1 | LP2  | Lemari Pakaian Plastik  |     34 | 2022-03-17 23:08:02 | lemari           |
|  5 |              2 |           1 | LP3  | Lemari Rak Buku         |     95 | 2022-03-17 23:08:02 | lemari           |
|  6 |              3 |           4 | BJ1  | Baju Anak               |   1000 | 2022-03-17 23:13:04 | baju             |
|  7 |              3 |           4 | BJ2  | Dress                   |    140 | 2022-03-17 23:13:04 | baju             |
|  8 |              3 |           4 | BJ3  | Kaos/T-shirt            |   1450 | 2022-03-17 23:13:04 | baju             |
+----+----------------+-------------+------+-------------------------+--------+---------------------+------------------+
6 rows in set (0.043 sec)

-- 5
MariaDB [alta_online_shop]> SELECT t.*, p.name AS produk_name, u.name AS user_name FROM transaction t JOIN produk p JOIN users u;

-- 6
MariaDB [alta_online_shop]> DELIMITER $$
MariaDB [alta_online_shop]> CREATE TRIGGER deleteTransdet_Trans
    -> BEFORE DELETE ON transaction FOR EACH ROW
    -> BEGIN
    -> DELETE FROM transaction_details WHERE transaction_id=OLD.id;
    -> END$$
Query OK, 0 rows affected (0.018 sec)

MariaDB [alta_online_shop]> DELIMITER ;
MariaDB [alta_online_shop]> show TRIGGERS;

-- 7
MariaDB [alta_online_shop]> DELIMITER $$
MariaDB [alta_online_shop]> CREATE TRIGGER updateTotal_qyt
    -> BEFORE DELETE ON transaction_details FOR EACH ROW
    -> BEGIN
    -> DECLARE v_transaction_id INT;
    -> DECLARE new_total_qyt INT;
    -> SET v_transaction_id = OLD.transaction_id;
    -> SET new_total_qyt = (SELECT SUM(qyt) FROM transaction_details WHERE transaction_id = v_transaction_id) - OLD.qyt;
    -> UPDATE transaction SET total_qyt = new_total_qyt WHERE id = v_transaction_id;
    -> END$$
Query OK, 0 rows affected (0.045 sec)

MariaDB [alta_online_shop]> DELIMITER ;
MariaDB [alta_online_shop]> show triggers;

-- 8
MariaDB [alta_online_shop]> select * from produk where id NOT IN (select produk_id from transaction_details);
Empty set (0.013 sec)