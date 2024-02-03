package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, name)
	} else {
		fmt.Fprintf(w, "Hello %s \n", name)
	}
}

func TestQueryParameter(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Adib", nil)

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)

}

func MultipleQueryParamete(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s \n", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Adib&last_name=sofyan", nil)

	MultipleQueryParamete(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}

func MultipleQueryParameterValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprintln(w, strings.Join(names, " "))
}

func TestMultipleQueryParameterValues(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=adib&name=hauzan&name=sofyan", nil)

	MultipleQueryParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
