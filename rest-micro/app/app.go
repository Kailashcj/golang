package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)
//Making a test change here to test my git pull request thing...
//App struct
type App struct {
	Router   *mux.Router
	Database *sql.DB
}

//SetupRouter setting GET and POST methods
func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/endpoint/{id}").
		HandlerFunc(app.getFunction)

	app.Router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(app.postFunction)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}
	dbdata := &DbData{}
	err := app.Database.QueryRow("SELECT id, date, name FROM `test` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Date, &dbdata.Name)
	if err != nil {
		log.Fatal("Database SELECT failed")
	}
	log.Println("you fetched a thing!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}

func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	e := app.Database.Ping()
	if e != nil {
		fmt.Println("ping failed")
	}
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('Sarthak')")
	if err != nil {
		log.Fatal("Database insert failed")
	}
	log.Println("You called a thing!")
	w.WriteHeader(http.StatusOK)
}

//DbData id,date and name fields
type DbData struct {
	ID   int       `json:"id"` //mapping the fields to Json
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}
