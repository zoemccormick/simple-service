package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/ingress", Ingress)
	http.HandleFunc("/egress", Egress)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/code/{code}", code)
	http.ListenAndServe(":8080", nil)
}

// HomePage returns default at root
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached simple-service, try /ingress or /egress")
}

// Ingress response with `hello`
func Ingress(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ingress request to simple-service successful!")
}

// Egress takes environment variable and sends egress request
func Egress(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Redirect resp to: %s", os.Getenv("EGRESS_ROUTE"))
	resp, err := http.Get(os.Getenv("EGRESS_ROUTE"))
	if err != nil {
		fmt.Fprintf(w, "Egress attempt failed.")
	}

	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Fprint(w, responseString)

}

func code(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c, ok := vars["code"]
	if !ok {
		fmt.Println("code is missing in parameters")
	}
	code, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println("code in parameters must be numeric")
	}

	text := http.StatusText(code)
	if text == "" {
		text = "Unknown code"
	}

	w.WriteHeader(code)
	w.Write([]byte(text))
}
