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

## GET
List Personalities

```
curl -v -X GET http://localhost:8000/api/personalities -H 'Content-Type: application/json'
```

Fetch Personality by id

```
curl -v -X GET http://localhost:8000/api/personalities/1 -H 'Content-Type: application/json'
```

## POST 

Create Personality

```
curl -v -X POST http://localhost:8000/api/personalities -H 'Content-Type: application/json' -d '{"name":<personality_name>,"history":<personality_history>}'
```