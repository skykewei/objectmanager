// user.go
package routers

import (
	"github.com/gorilla/mux"
)

func SetUserRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
}
