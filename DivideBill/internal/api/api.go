package api

import (
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
		service: service.New(),
	}
	
	EndPoints(api)

	return &api
}

func EndPoints(api API) {
	api.router.HandleFunc("/divide-bill", api.GetDivideBill()).Methods(http.MethodGet)
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) GetDivideBill() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hash := r.Header.Get("Hash-cheque")

		msg := api.service.Divide(hash)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(msg))
	})
}