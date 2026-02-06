package main

import (
	"fmt"
	"net/http"

	"de.marcules.httptesting/routes"
	"de.marcules.httptesting/routes/upload"
	"de.marcules.httptesting/routes/users"
)

var PORT = "8080"

func main() {
	mux := http.NewServeMux()
	registerRoutes(mux)
	server := http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	fmt.Printf("listening on http://localhost:%s", PORT)
	err := server.ListenAndServe()
	if err != nil {
		println(err.Error())
		return
	}
}

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", routes.RootHandler)
	mux.HandleFunc("GET /users", users.GetUsers)
	mux.HandleFunc("PUT /upload", upload.FileUploadHandler)
}
