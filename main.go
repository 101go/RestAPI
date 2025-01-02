package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	// variable name, type, and json key
	// json key is used to convert the struct to json
	ID         string `json:"id"`
	Item       string `json:"item"`
	Compeleted bool   `json:"compeleted"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go", Compeleted: false},
	{ID: "2", Item: "Build REST API", Compeleted: false},
	{ID: "3", Item: "Learn Docker", Compeleted: false},
}

func getTodos(context *gin.Context) {
	// context will contain information on the incoming http request

	// .IndentedJSON() will convert the todos variables above to json
	// http.StatusOK is the status code 200 which means the request was successful and it will return 200
	context.IndentedJSON(http.StatusOK, todos)
}

// run the app by running `go run main.go` in the terminal
func main() {
	router := gin.Default()        // create a new gin router, the router is our server
	router.GET("/todos", getTodos) // create a new route that listens to GET requests on /todos and calls the getTodos function when a request is made
	router.Run("localhost:9999")   // run the server on localhost:9999
}
