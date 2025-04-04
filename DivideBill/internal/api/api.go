package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
}

func New() *API {
	api := API{
		router: mux.NewRouter(),
	}
	
	EndPoints(api)

	return &api
}

func EndPoints(api API) {
	api.router.HandleFunc("/", api.GetDivideBill()).Methods(http.MethodGet)
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) GetDivideBill() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {


		
	})
}