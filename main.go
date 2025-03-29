package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// adding the gin import here and then executing `go get .` from
// the command line will add gin to the module as a dependency.
// go will also resolve and download the dependency to satisfy the import

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// the `json` struct tags above will specify what the name of the field
// should be when the struct's contents are serialized to json.
// without the tag, the capitalized name of the field would be used in the json

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// album data is only hardcoded in memory here to simplify the project
// so that a database connection is not needed.

// *** MAIN ***

func main() {
	router := gin.Default()          // initialize a gin router
	router.GET("/albums", getAlbums) // map a url path and http GET method to a handler

	router.Run("localhost:8080") // attach router to an http.Server and start server
}

// *** REQUEST HANDLERS ***

// getAlbums responds with a list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// gin.Context is the "most important part of gin" and it contains
// request data, errors, can validate JSON, etc.
// https://pkg.go.dev/github.com/gin-gonic/gin#Context

// we are using Context.IndentedJSON to serialize data from the albums struct
// and include it in the response. the first arg passed is the status code
// which will be returned by the response.
