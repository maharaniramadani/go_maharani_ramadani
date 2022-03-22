> use alta_online_shop
switched to db alta_online_shop
> show collections
> db.createCollection('users')
{ "ok" : 1 }
> db.createCollection('products')
{ "ok" : 1 }
> db.createCollection('product_descriptions')
{ "ok" : 1 }
> db.createCollection('product_types')
{ "ok" : 1 }
> db.createCollection('operators')
{ "ok" : 1 }
> db.createCollection('payment_methods')
{ "ok" : 1 }
> db.createCollection('transactions')
{ "ok" : 1 }
> db.createCollection('transactions_details')
{ "ok" : 1 }
> show collections
operators
payment_methods
product_descriptions
product_types
products
transactions
transactions_details
users

// 1
db.operators.insertMany([
    {_id: 1, name: 'operator 1'},
    {_id: 2, name: 'operator 2'},
    {_id: 3, name: 'operator 3'},
    {_id: 4, name: 'operator 4'},
    {_id: 5, name: 'operator 5'},
])

// 2
db.product_types.insertMany([
    {_id: 1, name: 'skincare'},
    {_id: 2, name: 'lemari'},
    {_id: 3, name: 'baju'},
])

// 3
db.product.insertMany([
    {_id: 1, product_type_id: 1, operator_id: 3, code: 'SC1', name: 'Toner', status: 200},
    {_id: 2, product_type_id: 1, operator_id: 3, code: 'SC1', name: 'Serum', status: 215},
    {_id: 3, product_type_id: 2, operator_id: 1, code: 'LP1', name: 'Lemari Pakaian Portable', status: 20},
    {_id: 4, product_type_id: 2, operator_id: 1, code: 'LP2', name: 'Lemari Pakaian Plastik', status: 34},
    {_id: 5, product_type_id: 2, operator_id: 1, code: 'LP3', name: 'Lemari Rak Buku', status: 95},
    {_id: 6, product_type_id: 3, operator_id: 4, code: 'BJ1', name: 'Baju Anak', status: 1000},
    {_id: 7, product_type_id: 3, operator_id: 4, code: 'BJ2', name: 'Dress', status: 140},
    {_id: 8, product_type_id: 3, operator_id: 4, code: 'BJ3', name: 'Kaos/T-shirt', status: 1450},
])

// 4
db.product_descriptions.insertMany([
    {_id:1, produk_id: 1, descriptions: 'Toner merupakan produk perawatan kulit untuk menghapus sisa makeup,kotoran,minyak dan mencerahkan wajah'},
    {_id:2, produk_id: 2, descriptions: 'Serum merupakan produk perawatan kulit wajah untuk melembabkan, mencerahkan, hingga mengatasi kerutan'},
    {_id: 3, produk_id: 3, descriptions:'Lemari pakaian portable berbahan kain yang bisa dilipat'},
    {_id: 4, produk_id: 4, descriptions: 'Lemari pakaian plastik berkualitas tinggi ukuran 2x3 tingkat'},
    {_id: 5, produk_id: 5, descriptions: 'Lemari Buku, lemari portable serbaguna untuk merapikan buku, tempat majalah atau dokumen'},
    {_id: 6, produk_id: 6, descriptions: 'Baju anak usia 1-5 tahun bahan katun'},
    {_id: 7, produk_id: 7, descriptions: 'Dress cantik berkualitas premium'},
    {_id: 8, produk_id: 8, descriptions:'Kaos/Tshirt oblong berkualitas premium overzise'},
])

// 5
db.payment_methods.insertMany([
    {_id:1, name: 'altapay', status: 200},
    {_id: 2,name: 'E-Banking', status: 200},
    {_id: 3,name: 'Transfer ATM', status: 200},
])
// 6
db.users.insertMany([
    {_id:1,name: 'rani',status: 200, dob: Date('2001-11-17'),gender: 'F'},
    {_id:2, name:'nia', status: 200, dob: Date('2001-03-29'),gender: 'F'},
    {_id: 3,name: 'ahmad',status: 200, dob: Date('2001-02-02'), gender:'M'},
    {_id: 4,name: 'roy', status:200,dob: Date ('2001-12-09'),gender: 'M'},
    {_id:5, name:'andin', status:200,dob: Date ('2000-10-20'), gender:'F'},
])

// 7
db.transactions.insertMany([
    {_id:1,user_id: 1,payment_method_id :1, status: 'sukses',total_qty:2,total_price: 90000},
    {_id:2,user_id:1,payment_method_id :2,status:'sukses',total_qty:4,total_price:220000},
    {_id:3,user_id:1,payment_method_id :3,status:'sukses',total_qty:3,total_price:560000},
    {_id:4,user_id:2,payment_method_id :1,status:'tunda',total_qty:5,total_price:755000},
    {_id:5,user_id:2,payment_method_id :2,status:'sukses',total_qty:1,total_price:45000},
    {_id:6,user_id:2,payment_method_id :3,status:'batal',total_qty:3,total_price:260000},
    {_id:7,user_id:3,payment_method_id :1,status:'batal',total_qty:6,total_price:800000},
    {_id:8,user_id:3,payment_method_id :2,status:'sukses',total_qty:1,total_price:80000},
    {_id:9,user_id:3,payment_method_id :3,status:'sukses',total_qty:2,total_price:100000},
    {_id:10,user_id:4,payment_method_id :1,status:'tunda',total_qty:3,total_price:110000},
    {_id:11,user_id:4,payment_method_id :2,status:'sukses',total_qty:5,total_price:225000},
    {_id:12,user_id:4,payment_method_id :3,status:'tunda',total_qty:2,total_price:110000},
    {_id:13,user_id:5,payment_method_id :1,status:'sukses',total_qty:4,total_price:100000},
    {_id:14,user_id:5,payment_method_id :2,status:'sukses',total_qty:5,total_price:755000},
    {_id:15,user_id:5,payment_method_id :3,status:'tunda',total_qty:4,total_price:495000},
])

// 8
db.transactions_details.insertMany([
    {transactions_id: 1,produk_id: 1,status: 'selesai',qty: 2,price: 90000},
    {transactions_id:1,produk_id: 2, status:'dikirim', qty:1, price:45000},
    {transactions_id:1,produk_id: 3, status:'dikirim', qty:4,price: 220000},
    {transactions_id:2, produk_id:4, status:'selesai', qty:3, price:200000},
    {transactions_id:2, produk_id:5,status: 'selesai', qty:1,price: 110000},
    {transactions_id:2, produk_id:6, status:'selesai', qty:4,price: 590000},
    {transactions_id:3, produk_id:7, status:'selesai', qty:3,price: 900000},
    {transactions_id:3,produk_id: 8, status:'selesai', qty:5,price: 900000},
    {transactions_id:3, produk_id:1, status:'selesai', qty:1,price: 100000},
    {transactions_id:4, produk_id:2, status:'selesai', qty:3,price: 500000},
    {transactions_id:4,produk_id: 3, status:'selesai', qty:2,price: 80000},
    {transactions_id:4, produk_id:4,status: 'dikirim', qty:2,price: 900000},
    {transactions_id:5,produk_id: 5,status: 'selesai', qty:2,price: 300000},
    {transactions_id:5, produk_id:6, status:'selesai', qty:5,price: 350000},
    {transactions_id:5, produk_id:7,status: 'selesai',qty: 3,price: 650000},
    {transactions_id:6, produk_id:8, status:'selesai', qty:2,price: 300000},
    {transactions_id:6, produk_id:1, status:'selesai', qty:2,price: 110000},
    {transactions_id:6, produk_id:2,status: 'selesai', qty:2,price: 120000},
    {transactions_id:7, produk_id:3, status:'dikirim', qty:4, price:900000},
    {transactions_id:7, produk_id:4, status:'dikirim', qty:5,price: 1500000},
    {transactions_id:7, produk_id:5,status: 'selesai', qty:1,price: 450000},
    {transactions_id:8, produk_id:6,status: 'selesai', qty:12,price: 540000},
    {transactions_id:8, produk_id:7,status: 'selesai', qty:12,price: 540000},
    {transactions_id:8, produk_id:8,status: 'selesai', qty:24,price: 2400000},
    {transactions_id:9, produk_id:1,status: 'selesai',qty: 4, price:220000},
    {transactions_id:9, produk_id:2, status:'selesai', qty:4, price:220000},
    {transactions_id:9, produk_id:3, status:'selesai', qty:2, price:900000},
    {transactions_id:10, produk_id:4,status: 'dikirim',qty: 2,price: 1000000},
    {transactions_id:10, produk_id:5, status:'selesai',qty: 1,price: 450000},
    {transactions_id:10, produk_id:6, status:'selesai',qty: 1,price: 500000},
    {transactions_id:11, produk_id:7, status:'selesai',qty: 1,price: 150000},
    {transactions_id:11, produk_id:8, status:'selesai', qty:5,price: 450000},
    {transactions_id:11, produk_id:1, status:'selesai', qty:3,price: 450000},
    {transactions_id:12, produk_id:2, status:'selesai', qty:5,price: 325000},
    {transactions_id:12, produk_id:3, status:'selesai', qty:1,price: 350000},
    {transactions_id:12, produk_id:4, status:'dikirim', qty:5,price: 2250000},
    {transactions_id:13, produk_id:5, status:'dikirim', qty:8, price:2400000},
    {transactions_id:13, produk_id:6, status:'selesai', qty:12,price: 540000},
    {transactions_id:13, produk_id:7, status:'selesai', qty:12,price: 700000},
    {transactions_id:13,produk_id: 8, status:'selesai', qty:5, price:450000},
    {transactions_id:14, produk_id:1, status:'selesai', qty:4,price: 180000},
    {transactions_id:14,produk_id: 2, status:'selesai', qty:8,price: 520000},
    {transactions_id:14, produk_id:3, status:'dikirim', qty:8,price: 2400000},
    {transactions_id:15, produk_id:4, status:'selesai', qty:1,price: 450000},
    {transactions_id:15, produk_id:5, status:'selesai', qty:1,price: 500000},
    {transactions_id:15, produk_id:6, status:'selesai', qty:24,price: 2500000},
])

// 9
db.users.find({gender:'M'})

{ "_id" : 3, "name" : "ahmad", "status" : 200, "dob" : "Tue Mar 22 2022 00:42:41 GMT-0700 (Pacific Daylight Time)", "gender" : "M" }
{ "_id" : 4, "name" : "roy", "status" : 200, "dob" : "Tue Mar 22 2022 00:42:41 GMT-0700 (Pacific Daylight Time)", "gender" : "M" }

// 10
db.product.find({_id:3})

{ "_id" : 3, "product_type_id" : 2, "operator_id" : 1, "code" : "LP1", "name" : "Lemari Pakaian Portable", "status" : 20 }

// 11
db.users.count({gender:'F'})
3

// 12
db.users.find().sort({name:1})
{ "_id" : 3, "name" : "ahmad", "status" : 200, "dob" : "Tue Mar 22 2022 00:42:41 GMT-0700 (Pacific Daylight Time)", "gender" : "M" }
{ "_id" : 5, "name" : "andin", "status" : 200, "dob" : "Tue Mar 22 2022 00:42:41 GMT-0700 (Pacific Daylight Time)", "gender" : "F" }
{ "_id" : 2, "name" : "nia", "status" : 200, "dob" : "Tue Mar 22 2022 00:42:41 GMT-0700 (Pacific Daylight Time)", "gender" : "F" }
{ "_id" : 1, "name" : "rani", "status" : 200, "dob" : "Tue Mar 22 2022 00:42:41 GMT-0700 (Pacific Daylight Time)", "gender" : "F" }
{ "_id" : 4, "name" : "roy", "status" : 200, "dob" : "Tue Mar 22 2022 00:42:41 GMT-0700 (Pacific Daylight Time)", "gender" : "M" }

// 13
db.product.find().limit(5)
{ "_id" : 1, "product_type_id" : 1, "operator_id" : 3, "code" : "SC1", "name" : "Toner", "status" : 200 }
{ "_id" : 2, "product_type_id" : 1, "operator_id" : 3, "code" : "SC1", "name" : "Serum", "status" : 215 }
{ "_id" : 3, "product_type_id" : 2, "operator_id" : 1, "code" : "LP1", "name" : "Lemari Pakaian Portable", "status" : 20 }
{ "_id" : 4, "product_type_id" : 2, "operator_id" : 1, "code" : "LP2", "name" : "Lemari Pakaian Plastik", "status" : 34 }
{ "_id" : 5, "product_type_id" : 2, "operator_id" : 1, "code" : "LP3", "name" : "Lemari Rak Buku", "status" : 95 }


// 14
db.product.updateOne({_id:1},{'$set':{name:'product dummy'}})

{ "acknowledged" : true, "matchedCount" : 1, "modifiedCount" : 1 }

// 15
db.transaction_details.updateMany({produk_id: 1}, {'$set': {qty: 3}})

{ "acknowledged" : true, "matchedCount" : 0, "modifiedCount" : 0 }

// 16
db.products.deleteOne({_id: 1})
{ "acknowledged" : true, "deletedCount" : 0 }
// 17
db.products.deleteMany({product_type_id: 1})
{ "acknowledged" : true, "deletedCount" : 0 }
