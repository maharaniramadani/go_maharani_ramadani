# 15 Introduction NoSQL and MongoDB

## Resume

Dalam materi ini, mempelajari:

1. Sejarah
2. No SQL Database
3. Tipe/Kategori

### Pengertian

Benefit Relational Database yaitu:

1. Dirancang untuk segala keperluan
2. SQL - Memiliki standar yang jelas
3. Memiliki banyak tool seperti (administrasi,reporting,framework)

Prinsip RDB adalah ACID:

1. Atomicity
2. Consistency
3. Isolation
4. Durabililty

NoSQL yaitu DBMS yang menyediakan mekanisme yang lebih fleksibel dibandingkan dengan RDBMS dan kelebihanya sendiri schema less, fast development, support big size file, support cluster.

Tipe/Kategori nosql

1. key/value (redis)
2. column family (cassandra)
3. graph (neo4j)
4. document based (mongodb)

### Penggunaan

jika membutuhkan skema fleksibel, acid tidak diperlukan, terdistribusi, data loggin, data sementara sebaiknya menggunakan NoSQL. jika untuk data finansial, data transaksi, business critical, acid compliant sebaiknya tidak menggunakan NoSQL.
