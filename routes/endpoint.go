package routes

import "net/http"

func Endpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
