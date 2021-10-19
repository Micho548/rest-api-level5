package postget

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method for get: "+r.Method)
	fmt.Println("Method used:", r.Method)
}

func Post(w http.ResponseWriter, r *http.Request) {
	ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "Method for post: "+r.Method)
	fmt.Println("Method used:", r.Method)
}
