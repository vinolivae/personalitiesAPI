package main

import (
	db "github.com/vinolivae/personalitiesAPI/database"
	r "github.com/vinolivae/personalitiesAPI/router"
)

func main() {
	db.DbConnect()
	r.LoadRoutes()
}
