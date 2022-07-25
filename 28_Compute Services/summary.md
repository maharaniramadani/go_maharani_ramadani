# 28 Compute Services

## Resume
Dalam materi ini, mempelajari:


1. Siklus Devops
2. Proses deployment


### Pengertian
System & Software Deployment


Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. Penyebarannya dapat melalui beragam cara tergantung dari jenis aplikasinya, aplikasi web & api ke server sedangkan aplikasi mobile ke Playstore/Appstore.


ada beberapa strategi: 
1. big-bang deployment strategy atau sering disebut replace deployment strategy,
2. rollout deployment strategy, 
3. blue/green deployment strategy, 
4. canary deployment strategy.


siklus devops 
1. perencanaan (plan) 
2. membuat program (code) 
3. build (build)
4. diuji coba (test) 
5. release version (release) 
6. deploy (deploy) 
7. beroperasi (operate) 
8. harus memonitoring (monitor)


proses deployment meliputi melakukan development dan mencobanya di localhost, lalu melakukan testing terlebih dahulu, lalu memeriksa user acceptance Ttesting agar kita bisa tahu apakah program yang kita kerjakan sudah sesuai dengan keinginan user, setelah UAT lolos maka kita bisa langsung ke tahap production.