package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	handler "github.com/autoremedy/go-function-sdk"
	"github.com/autoremedy/remedy-template/template/go-remedy/function"
	"github.com/prometheus/alertmanager/template"
)

func writeResponse(w http.ResponseWriter, msg string, err error) {
	log.Printf("%s: %v", msg, err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func makeRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			writeResponse(w, "expected request body", nil)
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeResponse(w, "failed to read body from request", err)
			return
		}

		var data template.Data
		err = json.Unmarshal(body, &data)
		if err != nil {
			writeResponse(w, "failed to unmarshal request body", err)
			return
		}

		req := handler.Request{
			Data:   data,
			Header: r.Header,
			Method: r.Method,
		}

		result, resultErr := function.Handle(req)

		if result.Header != nil {
			for k, v := range result.Header {
				w.Header()[k] = v
			}
		}

		if resultErr != nil {
			log.Print(resultErr)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if result.StatusCode == 0 {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(result.StatusCode)
			}
		}

		w.Write(result.Body)
	}
}

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8082),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	http.HandleFunc("/", makeRequestHandler())
	log.Fatal(s.ListenAndServe())
}
