package main

import (
	"net/http"

	"github.com/iamolegga/rebusexample/internal"
)

func main() {
	handler := internal.Build()
	http.Handle("/", handler)
	_ = http.ListenAndServe(":8080", nil)
}
