package main

import (
	"fmt"
	"log" //log package is used to print logs (messages) with timestamps
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	
	if r.Method != "GET"{
		http.Error(w,"Method is not supported",http.StatusNotFound) // Send a 404 error if the request method is not GET
		return
	}
	fmt.Fprintf(w,"Hello!") // Write "Hello!" to the response
}
func main(){
	fileserver := http.FileServer(http.Dir("./static")) // Create a file server that serves files from the "static" directory
	http.Handle("/",fileserver)// Register the file server to handle requests to the root path "/"
	http.HandleFunc("/hello",helloHandler) // Register the helloHandler function to handle requests to the "/hello" path
	http.HandleFunc("/form",formHandler) // Register the formHandler function to handle requests to the "/form" path

	fmt.Printf("Starting server at port 8080\n") // Print a message indicating that the server is starting
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)// Log and exit if the server fails to start
	}
}