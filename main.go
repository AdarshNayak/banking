package main

/*  Reference for more info about HTTP server concepts:
https://golangbyexample.com/basic-http-server-go/
https://betterprogramming.pub/overview-of-server-side-http-apis-in-go-44f052737e4b
*/

/* Execution Steps:
1. Execute the main.go in one terminal -> which would start server,
2. open 2nd terminal and call "curl http://localhost:8000/greet". This would print "Hello World!!"
*/

// Keep only the application logic in main
func main() {
	Start()
}
