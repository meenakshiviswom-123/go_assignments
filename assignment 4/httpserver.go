package main
import (
	"fmt"
	"net/http"
)
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler2)
	http.ListenAndServe(":8080", nil)
	}
	func handler (W http.ResponseWriter, r *http.Request){
	fmt.Fprintf(W, "Welcome to my home page\n")
	}
	func handler2 (W http.ResponseWriter, r *http.Request){
	fmt.Fprintf(W, "hello world\n")
	}