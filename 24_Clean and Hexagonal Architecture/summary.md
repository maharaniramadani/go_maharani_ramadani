# 23 Unit Test

## Resume

Dalam materi ini, mempelajari:

1. Clean Architecture
2. CA layer

### Pengertian

architecture yang bagus adalah architectur yang bisa terpisahkan secara layer, kita bisa membisahkan kebutuhan nya menggunakan layer dan membuat sebuah modular yang bisa di scalabel atau bisa di upragde bisa di perbaiki dan bisa di test di aplikasi tersebut.

ketika kita membuat suatu aplikasi kita pasti akan menyediakan tempat kosong untuk improvement, tanpa itu pasti akan sangat sulit.

hexagonal architecture terdiri dari API layer (http, message broker), domain layer (business logic), SPI layer (service provider)
kendala sebelum merancang clean architecture:

1. Independent of frameworks,
2. Testable,
3. Independent of UI,
4. Independent of database,
5. Independent of any external.

manfaat:
jika mempunyai standar structure akan lebih mudah untuk swicthing project, lebih cepat development,unit test selalu bisa berjalan dengan baik meskipun ada layer yang kurang baik.

clean architecture layer mempunyai 3 layer architecture:

- entities (business object),
- use case - domain layer (business logic),
- controller - presentation layer (handler),
- driver - data layer (data management).

Domain Driven Design (DDD):
pengembangan yang berorientasi pada domain (business logic)
CA adalah sebuah software architecture, sedangkan DDD adalah software design technique
