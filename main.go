package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, rootForm)
}

const rootForm = `
{
  "ruta":[
  	{
  		"lat":
  		"lng":
  	}
  ]
}
`
