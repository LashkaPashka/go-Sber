package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lashkapashka/divideBill/internal/service"
)

type API struct {
	router *mux.Router
	service *service.DivideService
}

func New() *API {
	api := API{
		router: mux.NewRouter(),
		service: service.NewDivideService(),
	}
	
	EndPoints(api)

	return &api
}

func EndPoints(api API) {
	api.router.HandleFunc("/split-position", api.GetDividePosition()).Methods(http.MethodPost)
	api.router.HandleFunc("/split-account", api.GetDivideAccount()).Methods(http.MethodPost)
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) GetDividePosition() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		
		var req map[string]string
		err := json.NewDecoder(r.Body).Decode(&req)
		
		if err != nil {
			panic(err)
		}
		
		msg := api.service.GetPosition(req)

		log.Println(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	})
}

func (api *API) GetDivideAccount() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		
		var req map[string]string
		err := json.NewDecoder(r.Body).Decode(&req)
		
		if err != nil {
			panic(err)
		}
		
		msg := api.service.GetAccount(req)
		log.Println(msg)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	})
}