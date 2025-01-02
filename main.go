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

func addTodos(context *gin.Context) {
	var newTodo todo // create a new todo variable
	// in go when declaring variables, you assign variable types after the variable name
	if err := context.BindJSON(&newTodo); // bind the incoming json request to the newTodo variable
	err != nil {                          // if there is an error, return.
		return
	}
	todos = append(todos, newTodo)                    // append the newTodo variable to the todos variable
	context.IndentedJSON(http.StatusCreated, newTodo) // return the newTodo variable as json with a status code of 201
}

func getTodoById(context *gin.Context) {
	id := context.Param("id")    // get the id from the url
	for _, todo := range todos { // loop through the todos variable
		if todo.ID == id { // if the todo.ID is equal to the id from the url
			context.IndentedJSON(http.StatusOK, todo) // return the todo as json with a status code of 200
			return
		}

	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"}) // if the todo is not found, return a status code of 404 with a message
}

// patch request: updating an item that already exists

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id") // get the id from the url
	for index, todo := range todos {
		if todo.ID == id {
			todos[index].Compeleted = !todos[index].Compeleted
			context.IndentedJSON(http.StatusOK, todos[index])
			return
		}
	}
}

// run the app by running `go run main.go` in the terminal
func main() {
	router := gin.Default()               // create a new gin router, the router is our server
	router.POST("/todos", addTodos)       // create a new route that listens to POST requests on /todos and calls the addTodos function when a request is made
	router.GET("/todos", getTodos)        // create a new route that listens to GET requests on /todos and calls the getTodos function when a request is made
	router.GET("/todos/:id", getTodoById) // create a new route that listens to GET requests on /todos/:id and calls the getTodoById function when a request is made
	router.POST("/todos/:id/toggle", toggleTodoStatus)
	router.Run("localhost:9999") // run the server on localhost:9999
}
