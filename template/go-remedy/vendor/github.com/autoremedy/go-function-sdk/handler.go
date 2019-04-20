package handler

import (
	"net/http"

	"github.com/prometheus/alertmanager/template"
)

// Response of function call
type Response struct {

	// Body the body will be written back
	Body []byte

	// StatusCode needs to be populated with value such as http.StatusOK
	StatusCode int

	// Header is optional and contains any additional headers the function response should set
	Header http.Header
}

// Configuration options for a Request
type Configuration struct {
	Options map[string]string
}

// Request of remediation function call
type Request struct {
	Alert       template.Alert
	Config      Configuration
	Header      http.Header
	QueryString string
	Method      string
}

// FunctionHandler used for a serverless Go method invocation
type FunctionHandler interface {
	Handle(req Request) (Response, error)
}
