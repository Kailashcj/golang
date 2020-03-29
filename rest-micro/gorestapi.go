package main

import (
	"log"
	"net/http"

	"microservice.com/microservice/db"

	"microservice.com/microservice/app"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

/*func setupRouter(router *mux.Router) {
	router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(postFunction)
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed")
	}
	_, err = database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database insert failed")
	}
	log.Println("You called a thing!")
}
*/
func main() {
	/*
		router := mux.NewRouter().StrictSlash(true)
		setupRouter(router)
		log.Fatal(http.ListenAndServe(":8080", router))
	*/
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}
	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
