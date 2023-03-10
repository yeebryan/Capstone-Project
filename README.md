# OJT


1. Pre-Made playlist: choose cuisine > interval > payment 
2. Self select: browse or filter restaurants and items > add items to "playlist cart" > interval > payment

For mongoDB:
Create a .env file in server folder with these lines of code
    PORT=3000
    MONGODB_URL=mongodb+srv://admin:Password@cluster0.pm6xbuy.mongodb.net/test
Replace admin with your username
Replace Password with your password

To run:
Run main.go in server first
cd to "/ojt/server" and type command "go run main.go"
Then cd to "/ojt/shop" and type command "npm start"