package healthresponse

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status int
	Body any
}

func OkResponse() Response {
	return Response{
	Status: http.StatusOK,
	Body: gin.H{
		"Status": "OK",
	}}}


func Health() chan Response {
	// Create a simple HTTP server for Kubernetes health checks
	return GinResponse("/health", OkResponse())
}

func GinResponse(path string, response Response) chan Response {
	c := make(chan Response)
	
	router := gin.Default()
	router.GET(path, func(c *gin.Context) {
		c.JSON(response.Status, response.Body)
	})
	go router.Run()
	go func() {
		for {
			response = <-c
		}
	}()
	fmt.Println("healthresponse started:", path, response)
	return c
}


