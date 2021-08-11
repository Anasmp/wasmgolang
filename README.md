### Features

- Wasm implimentation using golang
- Api binding from wasm
- Button hide and show from wasm
- Fetch the top 10 stories from Hacker News and displays them as list
- The examples require at least Go version 1.12



## Usage
 
####  To build wasm file from go file

`GOARCH=wasm GOOS=js go build -o lib.wasm main.go `

####  To run go static server 

`go run server.go`

####  After a successfull compilation 

######  Browse to [http://localhost:8080](http://localhost:8080 "http://localhost:8080") 


