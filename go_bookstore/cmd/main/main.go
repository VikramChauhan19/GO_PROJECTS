package main

import (
	"log"
	"net/http"
	"github.com/vikramchauhan19/go_bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter() //create a empty new router
	routes.RegisterBookstoreRoutes(r) //register all the routes
	http.Handle("/",r) //This tells Go: “For all incoming requests, use r (the router) to handle them.”
	log.Fatal(http.ListenAndServe("localhost:9010",r))//Start the web server on port 9010, and if it crashes, print the error and stop the program.

}
