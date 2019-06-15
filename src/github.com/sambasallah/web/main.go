package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"github.com/gorilla/mux"
)

// functions to handle routes

func Index(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("index.html"))
	index.Execute(w,nil)
}

func Products(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Products page")
	fmt.Println("Products Hit")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"About page")
	fmt.Println("About Hit")
}

func SingleProductView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Single product")
	fmt.Println("Single Product View Hit")
}

func Categories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"All Categories")
	fmt.Println("All Categories Hit")
}

func SingleCategoryView(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Single Category")
	fmt.Println("Single Category Hit")
}


func main() {
	
	r := mux.NewRouter()
	// defining routes
	r.HandleFunc("/",Index)

	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	log.Fatal(http.ListenAndServe(":8000",r))
}