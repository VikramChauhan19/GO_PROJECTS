package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vikramchauhan19/go_bookstore/pkg/models"
	"github.com/vikramchauhan19/go_bookstore/pkg/util"
)


func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res,_:=json.Marshal(newBooks)//Marshal = struct to json
	w.Header().Set("Content-Type","application/json") // “Hey browser, I’m sending JSON, not plain text.”
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)//Get route variables
	bookId := vars["bookId"]//“Read the values written in the URL.” bookId come from r.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	ID,err := strconv.ParseInt(bookId,0,0)//Convert string to int64, 0->auto base detection, 0->bitSize of the int
	if err != nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_ := models.GetBookById(ID)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
//Names starting with a lowercase letter are NOT accessible outside the package
func CreateBook(w http.ResponseWriter,r *http.Request){
	book := &models.Book{}
	util.ParseBody(r,book) //json to struct and move data to  book variable
	B:= book.CreateBook() // call createBook method of Book struct and created a entry in db
	res,_:=json.Marshal(B) //struct to json
	w.WriteHeader(http.StatusCreated) //201
	w.Header().Set("Content-Type","application/json")
	w.Write(res) 
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type","application/json") // “Hey browser, I’m sending JSON, not plain text.”
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){

	updateBook := &models.Book{}
	util.ParseBody(r,updateBook)

	vars:= mux.Vars(r)
	bookId := vars["bookId"]
	ID,err := strconv.ParseInt(bookId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}

	bookDetails,db := models.GetBookById(ID) // return pointer to Book and gorm.DB instance
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(bookDetails) //Changing a struct in Go does NOT automatically update the database. so we have to udate in db also
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}