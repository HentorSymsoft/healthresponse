package healthresponse

import (
	"testing"
	"time"
	"net/http"
)

func TestHealth(t *testing.T) {
	c := Health()

	time.Sleep(2 * time.Second)

	resp := Response{Status: http.StatusInternalServerError, Body: M{"error": "Error reason"}}

	c <- resp

	time.Sleep(2 * time.Second)

	resp = OkResponse()

	c <- resp

	time.Sleep(2 * time.Second)

}
