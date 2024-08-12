package server

import (
	"net/http"      // provides  http client and server implementations
	"text/template" // handle html templates, allows you to dynamically insert data into html files before sending them to the client

	function "server/ascii" // logic to generate ascii art
)

type ascii struct { // data structure that will be passed into the templates
	AsciiArt string // store  generated ascii art
	Error    string // store error message if any
	Output   string // store the original text input
}

// Handler for the homepage.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // ensures that only the get requests are accepted, else it returns a bad request error 400
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}

	home, err := template.New("index.html").ParseFiles("./template/index.html") // loads the index.html file form the template directory
	if err != nil {
		ErrorPageHandler(w, http.StatusInternalServerError, "500 - Internal Server Error") // return error if template cannot be loaded
		return
	}
	w.WriteHeader(http.StatusOK) // set  the response status code to 200 ok indicating a successful request, called before writing response body

	home.Execute(w, nil) // renders the index.html template and writes it to the http.ResponseWriter which sends it to the client browser
}

// Handler for the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}
	http.ServeFile(w, r, "./template/about.html") // for serving static files.
	// automatically returns a 200 ok status if the file is found and served successfully

	// tmpl, err := template.New("about.html").ParseFiles("./template/about.html")
	// if err != nil {
	// 	ErrorPageHandler(w, http.StatusInternalServerError, "400 Bad Request")
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)

	// tmpl.Execute(w, nil)
}

// Handler for the instruction page.
func InstructionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
	}

	http.ServeFile(w, r, "./template/instructions.html")

	// 	tmpl, err := template.New("instructions.html").ParseFiles("./template/instructions.html")
	// 	if err != nil {
	// 		ErrorPageHandler(w, http.StatusBadRequest, "400 Bad Request")
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusOK)

	// 	tmpl.Execute(w, nil)
}

// Handler for the error page.
func ErrorPageHandler(w http.ResponseWriter, statusCode int, message string) { // handles displaying error pages
	tmpl, err := template.New("error.html").ParseFiles("./template/error.html") // loads error.html template and populates it with an error message
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)     // sends specified http status code to client
	data := ascii{Error: message} // parses the error message to the template
	tmpl.Execute(w, data)         // renders html templates generate dynamic content by filling in placeholders within HTML files or other text format
}

// Handler used for obtaining ascii art
func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // ensures that only post methods are accepted
		ErrorPageHandler(w, http.StatusBadRequest, "400 - Bad Request")
		return
	}

	r.ParseForm()                // parses the form sent by the client from the post method
	text := r.FormValue("input") // retrieves the text input from the form

	banner := r.FormValue("bannerfile")            // gets the selected banner type from the form
	ascii_Art, err := function.Input(text, banner) // calls the function that generates ascii art using the text and banner type
	// these 2 values will be rendered in the ascii art template

	if text == "" || banner == "" { // error handling if the inputs are empty
		ErrorPageHandler(w, http.StatusBadRequest, "400- Bad Request")
		return
	}

	banners := []string{"standard", "thinkertoy", "shadow"}
	for i := range banners {
		if banner != banners[i] && i == len(banners)-1 {
			ErrorPageHandler(w, http.StatusBadRequest, "400 - Bad Request")
			return
		} else if banner == banners[i] {
			break
		}
	}
	if err != nil { // handle file not found error
		if err.Error() == "file not found" {

			ErrorPageHandler(w, http.StatusNotFound, "404 - Not Found")
			return
		}
		// all other errors are internal server errors
		ErrorPageHandler(w, http.StatusInternalServerError, "500 - Internal Server Error")
		return
	}

	data := ascii{AsciiArt: ascii_Art, Output: text} // data to be parsed into the ascii art output

	tmpl := template.Must(template.ParseFiles("template/ascii.html"))
	err1 := tmpl.Execute(w, data) // write data into the html template
	if err1 != nil {
		ErrorPageHandler(w, http.StatusBadRequest, "404 - Not Found")
		return
	}
}
