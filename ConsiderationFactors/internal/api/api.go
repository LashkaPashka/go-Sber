package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lashkapshka/go-Sber/internal/service"
)

type API struct {
	router *mux.Router
	service *service.FactorsService
}

func New() *API {
	api := &API{
		router: mux.NewRouter(),
		service: service.New(),
	}
	api.EndPoints()
	return api
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}


func (api *API) EndPoints() {
	api.router.HandleFunc("/using-factors", api.Apply()).Methods(http.MethodGet)
}


func (api *API) Apply() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		w.Write([]byte("Страница пользователя"))
	})
}

