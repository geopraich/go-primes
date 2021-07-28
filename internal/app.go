package internal

import (
	"encoding/json"
	"fmt"
	"github.com/geopraich/go-primes/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialise() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/api/v1/primes/{initial}", app.GetPrimes).Methods(http.MethodGet)
}

func (app *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, app.Router))
}

func (app *App) GetPrimes(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	input, err := strconv.Atoi(params["initial"])
	if err != nil {
		fmt.Errorf("invalid value %s", err)
	}

	primes, err := service.GeneratePrimes(input)
	if err != nil {
		fmt.Errorf("invalid value %s", err)
	}

	response, err := json.Marshal(primes)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write([]byte(response))
}
