package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/adamo57/adamo.io/server"
	"github.com/gorilla/mux"
)

var (
	// Info for logging information.
	Info *log.Logger

	// Warning for logging warnings.
	Warning *log.Logger

	// Error for logging errors.
	Error *log.Logger
)

// Init sets up the server logging when program starts.
func Init(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Init(os.Stdout, os.Stdout, os.Stderr)
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandler)

	s := server.New(r, ":0")

	Info.Println("server running on port :0")

	err := s.ListenAndServe()
	if err != nil {
		Error.Fatal("cannot start the server.")
	}
}

// IndexHandler handles the index page.
// TODO: Move Handlers to a new package.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles(path.Join("templates", "index.html"))
	if err != nil {
		http.Error(w, err.Error(), 500)
		Error.Printf("error parsing the template %v", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		Error.Printf("error executing the template %v", err)
	}
}
