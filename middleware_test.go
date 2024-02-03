package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Executed Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Executed Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error %s", err)
		}
	}()

	errorHandler.Handler.ServeHTTP(writer, request)
}

//func (handler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	defer func() {
//		err := recover()
//		if err != nil {
//			fmt.Println("Terjadi Error")
//			writer.WriteHeader(http.StatusInternalServerError)
//			fmt.Fprintf(writer, "Error %s \n", err)
//		}
//	}()
//	handler.Handler.ServeHTTP(writer, request)
//}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Middleware")
	})

	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("Upps")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
