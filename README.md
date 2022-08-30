# healthresponse

Simple http responder using gin-gonic spawned of as a goroutine.
Listens for changed response on channel returned.

Intented use: simple kubernetes healthcheck endpoint

Example:

````
c := healthresponse.Health()

time.Sleep(5 * time.Second)

resp := healthresponse.Response{Status: 500, Body: map[string]string{"error": "Error reason"}}

c <- resp
````
