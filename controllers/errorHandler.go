package controllers

import (
	"fmt"
	"net/http"
	"os"
)

func errHandler(w http.ResponseWriter, err error) {
	if os.Getenv("DEBUG") == "1" {
		fmt.Println(err)
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
