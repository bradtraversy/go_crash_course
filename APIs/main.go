package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"errors"
)

type book struct {
	ID       string	`json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book {
	{ ID:"1", Title:"In Search of Lost Time", Author:"M. Proust", Quantity:1 },
	{ ID:"2", Title:"The Great Gatsby", Author:"F. Scott Fitzgerald", Quantity:6 },
	{ ID:"3", Title:"War and Peace", Author:"Leo Tolstoy", Quantity:5 },
}

func getBooks( c *gin.Context ) {
	c.IndentedJSON( http.StatusOK, books )
}

func createBook( c *gin.Context ) {
	var newBook book
	if err := c.BindJSON( &newBook ); err != nil {
		return
	}
	books = append( books, newBook )
	c.IndentedJSON( http.StatusCreated, newBook )
}

func main() {
	router := gin.Default()
	router.GET( "/books", getBooks )
	router.POST( "/books", createBook )
	router.Run( "localhost:8080" )
}
