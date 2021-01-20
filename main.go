    package main

import (
	"fmt"
	"log"
	"net/http"
   "encoding/json"
	"github.com/gorilla/mux"
	"os/exec"
	"io/ioutil"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get-article", getArticle)
	log.Fatal(http.ListenAndServe(":8081", router))
	fmt.Fprintf(w, "Hello")
}

type JSONOutput struct {
	Year         string `json:"Year"`
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	reqBody,err:= ioutil.ReadAll(r.Body)
	var local JSONOutput
	json.Unmarshal(reqBody, &local)
	cmd := exec.Command("ruby", "ny_times.rb", local.Year)
	output,err := cmd.Output()
	if (err != nil) {
		fmt.Fprintf(w, "Anarchy!!!")
	}
	json.NewEncoder(w).Encode(string(output))

}
