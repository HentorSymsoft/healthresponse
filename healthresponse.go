package healthresponse

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	Status int
	Body   any
}

// M is a shortcut for map[string]any
type M map[string]any

func OkResponse() Response {
	return Response{
		Status: http.StatusOK,
		Body: gin.H{
			"Status": "OK",
		}}
}

// Create a simple HTTP server for Kubernetes health checks
func Health() chan Response {
	router := gin.Default();
	response := GinResponse(router, "/health", OkResponse());
	go router.Run()
	return response
}

// Create a simple HTTP server for Kubernetes health checks
func HealthWithRouter(router *gin.Engine) chan Response {
	response := GinResponse(router, "/health", OkResponse());
	return response
}

func GinResponse(router *gin.Engine, path string, response Response) chan Response {
	c := make(chan Response)

	router.GET(path, func(c *gin.Context) {
		c.JSON(response.Status, response.Body)
	})

	go func() {
		for {
			response = <-c
			log.Println("healthresponse got new response:", path, response)
		}
	}()

	log.Println("healthresponse started:", path, response)
	return c
}