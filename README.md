# API Description
This is a simple API written in Go. It returns data is JSON format

# HOW TO SETUP YOUR DATABASE
>Creeate a database with any name your like.
>Create 2 tables users and products
>>User Table Columns - (user_id,username,password,first_name,last_name,dob,address) - Products Table Columns(product_id,product_name,product_description,product_price,product_selling_price,sizes,colors,images)
>>>Replace your database parameters with the constant variables declared on top of the api.go file
# Endpoints
>/api/login - HTTP REQUEST METHOD IS POST
>>You send your username and password to this endpoint to receive an authorization token which you can use to make subsequent request to other endpoints that require authorization
>>>JSON data sample to send to receive an authorization token `{"username":"myusername","password":"mypass"}`

>/api/users - HTTP REQUEST METHOD IS GET
>>This endpoint returns a JSON data of all users in your user table

>/api/users/{id} - HTTP REQUEST METHOD IS GET
>> This endpoint returns a single user

>/api/products - HTTP REQUEST METHOD IS GET
>> This endpoint returns a JSON data of all products  in your products table

>/api/products/{id} - HTTP REQUEST METHOD IS GET
>> This endpoint returns a single product

>/api/create-user - HTTP REQUEST METHOD IS POST
>> You send a JSON data of the columns in the users table with corresponding values 

>/api/create-product - HTTP REQUEST METHOD IS POST
>> You send a JSON data of the columns in the products table with the corresponding values

>/api/delete-user/{id} - HTTP REQUEST METHOD IS DELETE
>> You pass in the id of the user you want to delete

>/api/delete-product/{id} - HTTP REQUEST METHOD IS DELETE
>> You pass in the id of the product you want to delete

>/api/update-user/{id} - HTTP REQUEST METHOD IS PUT
>> You pass in the id of the user you want to update

>/api/update-product/{id} - HTTP REQUEST METHOD IS PUT
>> You pass in the id of the product you want to update

