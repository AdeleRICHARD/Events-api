package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func Health(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	strDoc := []string{
		"ON",
	}
	fmt.Fprintf(w, strings.Join(strDoc, "\n"))
}
