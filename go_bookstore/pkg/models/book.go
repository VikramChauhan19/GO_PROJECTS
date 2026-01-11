package models
import (
	"github.com/jinzhu/gorm"
	"github.com/vikramchauhan19/go_bookstore/pkg/config"
)
var db *gorm.DB


type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){ // a funct which runs automatically when the package is imported
	config.Connect() // connect with db
	db = config.GetDB() //return the db instance
	db.AutoMigrate(&Book{}) //AutoMigrate = “Create or update table from struct”
}

func (b *Book) CreateBook() *Book{ //This function saves a Book object into the database and returns it.
	db.Create(b) //db.Create takes a pointer so it can modify your struct and write DB-generated values back into it.
	return b
}

func GetAllBooks()[]Book{
	var Books []Book
	db.Find(&Books) //Fetch all book records from the database and store them into Books.
	return Books
}

func GetBookById(Id int64)(*Book,*gorm.DB){
	var getBook Book
	db.Where("ID=?",Id).Find(&getBook)//Fetch the book record with the specified ID from the database and store it into getBook.
	return &getBook,db
}
func DeleteBook(Id int64) Book{
	var book Book
	db.Where("ID=?",Id).Find(&book)
	db.Delete(&book)
	return book
}

