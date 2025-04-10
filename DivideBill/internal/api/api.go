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
	api.router.HandleFunc("/divide-bill", api.GetDivideBill()).Methods(http.MethodPost)
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) GetDivideBill() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {		
		// var req map[string]string
		// err := json.NewDecoder(r.Body).Decode(&req)
		
		// if err != nil {
		// 	panic(err)
		// }
		
		// fmt.Println(req)
		// msg := api.service.Divide(req)
		// fmt.Println(msg)

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte(msg))
	})
}