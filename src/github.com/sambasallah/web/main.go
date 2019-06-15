package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

// functions to handle routes

func Index(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles("static/index.html"))
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
	
	// defining routes
	http.HandleFunc("/",Index)
	http.HandleFunc("/about",About)
	http.HandleFunc("/products",Products)
	http.HandleFunc("/products/{id}",SingleProductView)
	http.HandleFunc("/categories",Categories)
	http.HandleFunc("/categories/{id}",SingleCategoryView)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))

	log.Fatal(http.ListenAndServe(":8000",nil))
}