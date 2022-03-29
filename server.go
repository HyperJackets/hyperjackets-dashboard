package main
// designates the main package -> entry point of the application
// importing the required packages
import (
    "fmt"
    "log"
    "net/http"
)

// uses the *http.Request object to check if the requested path and type is correct
func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    fmt.Fprintf(w, "Hello")
}
func main() {
    // passing handlers to the server so it knows how to respond to incoming requests and which requests to accept
    // Sends a hello! message as we pass on the response to ResponseWriter
    http.HandleFunc("/hello", helloHandler)

    fmt.Printf("Starting server at port 8080\n")
    // Starts server and specifies port, second parameter is for HTTP/2 configuration
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
