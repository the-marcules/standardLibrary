package routes

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World!"))
	if err != nil {
		println(err.Error())
		return
	}
}
