package main

import (
	"fmt"
	"net/http"
	mux "github.com/gorilla/mux"
	"io/ioutil"
	"log"
)

func api(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req,_ := http.NewRequest("GET","http://localhost:8000/api/products", nil)

	res,err := client.Do(req)

	if err != nil {
        fmt.Fprintf(w, "Error: %s", err.Error())
	}
	
	body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, string(body))


}

func main() {

	router := mux.NewRouter();
	router.HandleFunc("/",api)
	log.Fatal(http.ListenAndServe(":8001",router))
}