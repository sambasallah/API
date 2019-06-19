package main

import (
	_"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"
	"time"
	"os"
)

func Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API up and running"))
}


var Secret = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("High Secret"))
})

func main() {
	router := mux.NewRouter();

	router.HandleFunc("/status", Status).Methods("GET")
	router.Handle("/secret", JWTMiddleware.Handler(Secret)).Methods("GET")
	router.Handle("/get-token",GetTokenHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000",handlers.LoggingHandler(os.Stdout,router)))
}

// code to generate json web token

var secretKey = []byte("secret")


var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// generating json web token
	token := jwt.New(jwt.SigningMethodHS256)

	// claiming ownership of the token
	claim := token.Claims.(jwt.MapClaims)

	claim["admin"] = true
	claim["name"] = "samba sallah"
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// sign or token with our secret key
	signedToken, _ := token.SignedString(secretKey)

	// write the token to the browser
	w.Write([]byte(signedToken))
})

var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	},
	SigningMethod:jwt.SigningMethodHS256,
})