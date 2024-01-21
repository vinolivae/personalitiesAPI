# personalitiesAPI

My first rest api with golang!

I'm using two libraries

[gorm](https://gorm.io/): An ORM for golang;

[gorilla mux](https://pkg.go.dev/github.com/gorilla/mux): An HTTP handler for golang.


# Run

- docker-compose up -d
- go run main.go
- [access localhost:8000](http://localhost:8000/)

PS..: The database has already been populated by docker-database-init.db file with two records.

# Available Routes

`GET "/"`

`GET "/api/personalities"`

`GET "/api/personalities/{id}"`