// AppObj.go

package swagger

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// Some comment here
	_ "github.com/lib/pq"
)

// App app
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// AppObj AppObj
var AppObj App

// Initialize Initialize
func Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	AppObj.DB, err = sql.Open("postgres", connectionString) //"user=postgres password=1 dbname=orderlunch sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	AppObj.Router = NewRouter()
}

// Run Run
func Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, AppObj.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}
