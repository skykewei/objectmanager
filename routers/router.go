package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Routes for the user entity
	router = SetUserRouters(router)
	//Routes for the logicalobject entity
	router = SetLogicalObjectRoutes(router)
	//Routes for the
	return router
}
