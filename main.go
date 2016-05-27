package main

import (
	"log"
	"net/http"

	"objectmanager/common"
	"objectmanager/routers"

	"github.com/codegangsta/negroni"
)

//Entry point of the program

func main() {
	//Calls startup logic
	common.StartUp()
	//Get the mux router object
	router := routers.InitRoutes()
	//Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
}
