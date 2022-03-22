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