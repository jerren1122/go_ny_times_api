    package main

import (
	"log"
	"net/http"
   "encoding/json"
	"github.com/gorilla/mux"
	"os/exec"
	"io/ioutil"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get-article", getArticle)
	log.Fatal(http.ListenAndServe(":8081", router))
}

type JSONOutput struct {
	Year         string `json:"Year"`
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	reqBody,_:= ioutil.ReadAll(r.Body)
	var local JSONOutput
	json.Unmarshal(reqBody, &local)
	cmd := exec.Command("ruby", "ny_times.rb", local.Year)
	output,_ := cmd.Output()
	json.NewEncoder(w).Encode(string(output))
}
