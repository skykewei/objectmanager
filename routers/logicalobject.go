package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/shijuvar/go-web/taskmanager/common"
	"github.com/shijuvar/go-web/taskmanager/controllers"
)

func SetLogicalObjectRoutes(router *mux.Router) *mux.Router {
	logicalObjectRouter := mux.NewRouter()
	logicalObjectRouter.HandleFunc("/logicalobjects", controllers.CreateLogicalObject).Method("POST")
	logicalObjectRouter.HandleFunc("/logicalobjects/{id}", controllers.UpdateLogicalObject).Method("PUT")
	logicalObjectRouter.HandleFunc("/logicalobjects/", controllers.GetLogicalObject).Method("GET")
	logicalObjectRouter.HandleFunc("/logicalobjects/{id}", controllers.GetLogicalObjectById).Method("GET")
	logicalObjectRouter.HandleFunc("/logicalobjects/users/{id}", controllers.GetLogicalObjectsByUser).Method("GET")
	logicalObjectRouter.HandleFunc("/logicalobjects/{id}", controllers.DeleteTask).Methods("DELETE")
	router.PathPrefix("/logicalobjects").Handler(negroni.New(
		negroni.HandlerFuc(common.Authorize),
		negroni.Wrap(logicalObjectRouter),
	))
	return router
}
