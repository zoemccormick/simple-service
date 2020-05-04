package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/ingress", Ingress)
	http.HandleFunc("/egress", Egress)
	http.HandleFunc("/", HomePage)
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
