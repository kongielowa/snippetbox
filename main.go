package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.

func home(w http.ResponseWriter, r *http.Request) {
	
		if  r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		
			
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.

func snippetView(w http.ResponseWriter, r *http.Request) {
	//Extract the value of the id parameter from the query string
	// and try to convert it to an integer using the strconv.Atoi() function
	// If it can't be converted to an integer or the value is less than 1,
	// we return a 404 page not found response
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWriter

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {

		//2 Use the Header().Set() method to add an 'Allow: POST' header
		//2 to the responde header map. The first parameter is the header name,
		//2 and the second parameter is the header value.
		w.Header().Set("Allow", http.MethodPost)
		
		//1 If it's not, use the w.WriteHeader() method to send a 405 status
		//1 code and the w.Write() method to write a "Method Not Allowed"
		//1 responde body. We then return from the function so that
		//1 the subsequent code is not excuted.

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))

}


func main() {
	
	//1 Use the http.NewServeMux() function to initialize a new servemux,
	//1 then reister the home function as the handler for the "/" URL pattern.

	//2 Register the two new handler functions and corresponding URL patterns with
	//2 the servemux, in exactly the same way that we did before.

	//1
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//2
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network adress to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message abd exit.
	// Note that any error returned by http.ListenAndServe() is always non-nil.

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
