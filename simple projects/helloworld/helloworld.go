package main

import (
	"fmt"
	"net/http"
)

func helohandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world !!")
}
func helloworld() {
	http.HandleFunc("/", helohandle)
	http.ListenAndServe(":8080", nil)

}
