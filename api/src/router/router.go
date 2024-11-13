package router

import "github.com/gorilla/mux"

func Create() *mux.Router {
	return mux.NewRouter()
}