package function

import (
	"fmt"
	"net/http"

	handler "github.com/autoremedy/go-function-sdk"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	message := fmt.Sprintf("Alert: %+v", req.Alert)

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
