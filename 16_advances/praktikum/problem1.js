> use alta_book
switched to db alta_book
> db.createCollection('books')
{ "ok" : 1 }
> db.createCollection('authors')
{ "ok" : 1 }
> db.createCollection('publishers')
{ "ok" : 1 }
> show collections
authors
books
publishers

db.books.insertMany([
    {_id: 1, title: 'Wawasan Pancasila',authorID: 1, publisherID: 1, price: 71200, status: {page: 324, weight: 300}, publishAt: new Date('2018-10-01'), category: ['social', 'politics']},
    {_id: 2, title: 'Mata Air Keteladanan', authorID: 1, publisherID: 2, price: 106260, status: {page: 672, weight: 650}, publishAt: new Date('2017-09-01'), category: ['social', 'politics']},
    {_id: 3, title: 'Revolusi Pancasila', authorID: 1, publisherID: 1, price: 54400, status: {page: 220, weight: 500}, publishAt: new Date('201505-01'), category: ['social', 'politics']},
    {_id: 4, title: 'Self Driving', authorID: 2, publisherID: 1, price: 58650, status: {page: 286, weight: 300}, publishAt: new Date('2018-05-01'), category: ['self-development']},
    {_id: 5, title: 'Self Disruption', authorID: 2, publisherID: 2, price: 83300, status: {page: 400, weight: 800},publishAt: new Date('2018-05-01'),category: ['self-development']},
  ])

db.authors.insertMany([
  {_id: 1, firstName: 'Yudi', lastName: 'Latif'},
  {_id: 2, firstName: 'Rhenald', lastName: 'Kasali'},
])
{ "acknowledged" : true, "insertedIds" : [ 1, 2 ] }

db.publishers.insertMany([
  {_id: 1, publisherName: 'Expose'},
  {_id: 2, publisherName: 'Mizan'},
])

db.publishers.insertMany([
  {_id: 1, publisherName: 'Expose'},
  {_id: 2, publisherName: 'Mizan'},
])

// 1
db.books.find({authorID:{'$in':[1,2]}})

// 2
db.books.find({authorID:1},{projection:{title:1,price:1}})

// 3
db.books.aggregate([
  {'$match':{authorID:2}},
  {'$group':{totalPages:{'$sum':'$status.page'},_id:null}},
])

// 4
db.authors.aggregate([
  {'$lookup':{from: 'books',localField:'_id',foreignField: 'authorID',as:'book'}},
  { '$unwind': '$books' },
  { '$unwind': '$books.status' },
  { '$unwind': '$books.category' },
])


db.books.aggregate([
  {'$lookup':{from: 'authors',localField:'authorID',foreignField: '_id',as:'author'}},
  { '$unwind': '$author' },
])

// 5
db.books.aggregate([
  {'$lookup':{from: 'authors',localField:'authorID',foreignField: '_id',as:'author'}},
  { '$unwind': '$author' },
  {'$lookup':{from: 'publishers',localField:'publisherID',foreignField: '_id',as:'publisher'}},
  { '$unwind': '$publisher' },
])

// 6
db.authors.aggregate([
  { '$lookup': { from: 'books', localField: '_id', foreignField: 'authorID', as: 'books' } },
  { '$project': { _id: { '$concat': ['$firstName', ' ', '$lastName'] }, number_of_books: { '$size': '$books' }, list_of_book: '$books.title' } },
])

// 7
db.books.aggregate([
{ '$project': { title: 1, price: 1, promo: { '$cond': { if: { '$lt': ['$price', 60000] }, then: '1%', else: { '$cond': { if: { '$lt': ['$price', 90000] }, then: '2%', else: '3%' } } } } } },
])

// 8
db.books.aggregate([
  { '$lookup': { from: 'authors', localField: 'authorID', foreignField: '_id', as: 'author' } },
  { '$sort': { price: -1 } },
  { '$unwind': '$author' },
  { '$lookup': { from: 'publishers', localField: 'publisherID', foreignField: '_id', as: 'publisher' } },
  { '$unwind': '$publisher' },
  { '$project': { title: 1, price: 1, author: { '$concat': ['$author.firstName', ' ', '$author.lastName'] }, publisher: '$publisher.publisherName' } },
])
// 9
db.books.aggregate([
  { '$lookup': { from: 'publishers', localField: 'publisherID', foreignField: '_id', as: 'publisher' } },
  { '$unwind': '$publisher' },
  { '$project': { title: 1, price: 1, publisher: '$publisher.publisherName' } },
])

db.books.aggregate([
  { '$lookup': { from: 'publishers', localField: 'publisherID', foreignField: '_id', as: 'publisher' } },
  { '$match': { _id: { '$in': [3, 4] } } },
  { '$unwind': '$publisher' },
  { '$project': { title: 1, price: 1, publisher: '$publisher.publisherName' } },
])