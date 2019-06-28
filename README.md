# API Description
This is a simple API written in Go. It returns data is JSON format

# HOW TO SETUP YOUR DATABASE
> Create a database with any name your like.
- Create 2 tables users and products
- User Table Columns - (user_id,username,password,first_name,last_name,dob,address) 
- Products Table Columns - (product_id,product_name,product_description,product_price,product_selling_price,sizes,colors,images)
- Replace your database parameters with the constant variables declared on top of the api.go file
# Endpoints
`POST /api/login` 
- You send your username and password to this endpoint to receive an authorization token which you can use to make subsequent request to other endpoints that require authorization
- JSON data sample to send to receive an authorization token `{"username":"myusername","password":"mypass"}`

`GET /api/users` 
- This endpoint returns a JSON data of all users in your user table

`GET /api/users/{id}`
- This endpoint returns a single user

`GET /api/products`
- This endpoint returns a JSON data of all products  in your products table

`GET /api/products/{id}`  
- This endpoint returns a single product

`POST /api/create-user` 
- You send a JSON data of the columns in the users table with corresponding values 

`POST /api/create-product` 
- You send a JSON data of the columns in the products table with the corresponding values

`DELETE /api/delete-user/{id}`
- You pass in the id of the user you want to delete

`DELETE /api/delete-product/{id}` 
- You pass in the id of the product you want to delete

`PUT /api/update-user/{id}` 
- You pass in the id of the user you want to update

`PUT /api/update-product/{id}`
- You pass in the id of the product you want to update

