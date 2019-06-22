package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"log"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"compress/gzip"
	"io/ioutil"
	"github.com/gorilla/handlers"
	"os"
	"github.com/dgrijalva/jwt-go"
	"github.com/auth0/go-jwt-middleware"
	"time"
)

type Users struct {
	Id string `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	DOB string `json:"dob"`
	Address string `json:"address"`
}

type UsersPostData struct {
	Data Users
}

type Products struct {
	Id string `json:"id"`
	ProductName string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice string `json:"product_price"`
	ProductSellingPrice string `json:"product_selling_price"`
	ProductSizes string `json:"product_sizes"`
	ProductColors string `json:"product_colors"`
	ProductImages string `json:"product_images"`
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var user Users
	var allUsers []Users
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Accept-Charset","utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	db, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT user_id,username,password,first_name,last_name,dob,address FROM users")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err := results.Scan(&user.Id,&user.Username,&user.Password,&user.FirstName,&user.LastName,&user.DOB,&user.Address)
		if err != nil {
			panic(err.Error())
		}
		allUsers = append(allUsers, user)
	}
	
    // Gzip data
    gz := gzip.NewWriter(w)
    json.NewEncoder(gz).Encode(allUsers)
	gz.Close()
	
	fmt.Println("AllUsers Endpoint Hit")
	
}

func UserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Accept-Charset","utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	param := mux.Vars(r)
	id := param["id"]
	db, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	single, err := db.Query("SELECT * FROM users WHERE user_id=?",id)
	if err != nil {
		panic(err.Error())
	}

	var result Users

	for single.Next() {
		err := single.Scan(&result.Id,&result.Username,&result.Password,&result.FirstName,&result.LastName,&result.DOB,&result.Address)

		if err != nil {
			panic(err.Error())
		}

		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(result)
		gz.Close()
	}

	
	fmt.Println("UserById Endpoint Hit")


}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var data map[string]string
	
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	err := json.Unmarshal([]byte(reqBody),&data);

	if err != nil {
		panic(err.Error())
	}
	user_id := data["user_id"]
	username := data["username"]
	password := data["password"]
	first_name := data["first_name"]
	last_name := data["last_name"]
	dob := data["dob"]
	address := data["address"]

	conn, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")

	defer conn.Close()

	if err != nil {
		panic(err.Error())
	}

	 insert,err := conn.Query("INSERT INTO users (user_id,username,password,first_name,last_name,dob,address) VALUES (?,?,?,?,?,?,?)",user_id,username,password,first_name,last_name,dob,address)

	 if err != nil {
		 panic(err.Error())
	 }else {
		 fmt.Println("Data successfully inserted into the database")
	 }

	 fmt.Println("done")
	 insert.Close()
	

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	param := mux.Vars(r)
	id := param["id"]

	db, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("DELETE FROM users WHERE user_id = ?",id)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("User deleted successfully")
	}

	
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	id := param["id"]

	var data map[string]string

	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal([]byte(reqBody),&data)

	

	if err != nil {
		panic(err.Error())
	}

	// user_id := data["user_id"]
	username := data["username"]
	password := data["password"]
	first_name := data["first_name"]
	last_name := data["last_name"]
	dob := data["dob"]
	address := data["address"]


	db, err := sql.Open("mysql", "root:Y7enqxal!@/ebaaba")

	defer db.Close()

	if err != nil {
		fmt.Println("There was an error connecting to the database")
	}
	result, err := db.Query("UPDATE users SET username=?, password=?, first_name=?, last_name=?, dob=?, address=? WHERE user_id = ?",username,password,first_name,last_name,dob,address,id)
	defer result.Close()
	if err != nil {
		fmt.Println("There was an error in your query")
	}else {
		fmt.Println("User data has been successfully updated")
	}

	fmt.Println("Update Users Endpoint")
}


func AllProducts(w http.ResponseWriter, r *http.Request) {
	var product Products
	var allProducts []Products
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Accept-Charset","utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	db, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	results, err := db.Query("SELECT product_id,product_name,product_description,product_price,product_selling_price,sizes,colors,images FROM products")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err := results.Scan(&product.Id,&product.ProductName,&product.ProductDescription,&product.ProductPrice,&product.ProductSellingPrice,&product.ProductSizes,&product.ProductColors,&product.ProductImages)
		if err != nil {
			panic(err.Error())
		}
		allProducts = append(allProducts, product)
	}
	
	gz := gzip.NewWriter(w)
	json.NewEncoder(gz).Encode(allProducts)

	gz.Close()

	fmt.Println("AllProducts Endpoint Hit")
}


func ProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Accept-Charset","utf-8")
	w.Header().Set("Content-Encoding", "gzip")
	param := mux.Vars(r)
	id := param["id"]
	db, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	single, err := db.Query("SELECT * FROM products WHERE product_id=?",id)
	if err != nil {
		panic(err.Error())
	}

	var result Products

	for single.Next() {
		err := single.Scan(&result.Id,&result.ProductName,&result.ProductDescription,&result.ProductPrice,&result.ProductSellingPrice,&result.ProductSizes,&result.ProductColors,&result.ProductImages)

		if err != nil {
			panic(err.Error())
		}

		gz := gzip.NewWriter(w)
		json.NewEncoder(gz).Encode(result)
		gz.Close()
	}

	
	fmt.Println("ProductById Endpoint Hit")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	var data map[string]string
	
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	err := json.Unmarshal([]byte(reqBody),&data);

	if err != nil {
		panic(err.Error())
	}
	product_id := data["product_id"]
	product_name := data["product_name"]
	product_description := data["product_description"]
	product_price := data["product_price"]
	product_selling_price := data["product_selling_price"]
	sizes := data["sizes"]
	colors := data["colors"]
	images := data["images"]

	conn, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")

	defer conn.Close()

	if err != nil {
		panic(err.Error())
	}

	 insert,err := conn.Query("INSERT INTO products (product_id,product_name,product_description,product_price,product_selling_price,sizes,colors,images) VALUES (?,?,?,?,?,?,?,?)",product_id,product_name,product_description,product_price,product_selling_price,sizes,colors,images)

	 if err != nil {
		 panic(err.Error())
	 }else {
		 fmt.Println("Data successfully inserted into the database")
	 }

	 fmt.Println("done")
	 insert.Close()
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	param := mux.Vars(r)
	id := param["id"]

	db, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	result, err := db.Query("DELETE FROM products WHERE product_id = ?",id)
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("Product deleted successfully")
	}
}

var UpdateProduct = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	param := mux.Vars(r)
	id := param["id"]

	var data map[string]string

	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal([]byte(reqBody),&data)

	

	if err != nil {
		panic(err.Error())
	}

	product_name := data["product_name"]
	product_description := data["product_description"]
	colors := data["colors"]
	images := data["images"]
	product_price := data["product_price"]
	product_selling_price := data["product_selling_price"]
	sizes := data["sizes"]


	db, err := sql.Open("mysql", "root:Y7enqxal!@/ebaaba")

	defer db.Close()

	if err != nil {
		fmt.Println("There was an error connecting to the database")
	}
	result, err := db.Query("UPDATE products SET product_name=?, product_description=?, product_price=?, product_selling_price=?, sizes=?, colors=?, images=? WHERE user_id = ?",product_name,product_description,product_price,product_selling_price,sizes,colors,images,id)
	defer result.Close()
	if err != nil {
		fmt.Println("There was an error in your query")
	}else {
		fmt.Println("User data has been successfully updated")
	}

	fmt.Println("Update Products Endpoint")
})



func main() {
	// router instance
	router := mux.NewRouter();

	// get endpoints 
	router.HandleFunc("/api/users",AllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", UserById).Methods("GET")
	router.HandleFunc("/api/products",AllProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}",ProductById).Methods("GET")

	// post endpoints
	router.HandleFunc("/api/create-user", CreateUser).Methods("POST")
	router.HandleFunc("/api/create-product", CreateProduct).Methods("POST")

	// delete endpoints
	router.HandleFunc("/api/delete-user/{id}",DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/delete-product/{id}",DeleteProduct).Methods("DELETE")

	// update endpoints
	router.HandleFunc("/api/update-user/{id}", UpdateUser).Methods("PUT")
	router.Handle("/api/update-product/{id}", JWTMiddleware.Handler(UpdateProduct)).Methods("PUT")

	// login to get token
	router.HandleFunc("/api/login",Login).Methods("POST")

	// status 
	router.HandleFunc("/api/status",Status).Methods("GET")

	// test
	router.Handle("/secret",JWTMiddleware.Handler(Secret)).Methods("GET")


	
	log.Fatal(http.ListenAndServe(":9000",handlers.LoggingHandler(os.Stdout,router)))

}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}

var secretKey = []byte("secret-key")

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var data map[string]string

	reqBody, err := ioutil.ReadAll(r.Body);

	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal([]byte(reqBody),&data)


	username := data["username"]
	password := data["password"]

	db, err := sql.Open("mysql","root:Y7enqxal!@/ebaaba")
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	var count int = 0
	row, _ := db.Query("SELECT username,password FROM users WHERE username=? AND password=?",username,password)
	
	for row.Next() {
		count++
	}

	if count == 1 {

	// generating json web token
	token := jwt.New(jwt.SigningMethodHS256)

	// claiming ownership of the token
	claim := token.Claims.(jwt.MapClaims)

	claim["admin"] = true
	claim["name"] = "samba sallah"
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// sign or token with our secret key
	signedToken, _ := token.SignedString(secretKey)

	correct := map[string]string{"Authorization Token":signedToken}
	json_correct, _ := json.Marshal(correct)
	w.Write([]byte(json_correct))

	}else {
		wrong := `{"error":["Username & Password do not match"]}`
		// json_wrong, _ := json.Marshal(wrong)
		w.Write([]byte(wrong))
	}	

}


var Secret = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a secret data"))
})


var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// generating json web token

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