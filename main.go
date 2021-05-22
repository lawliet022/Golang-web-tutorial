package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// "/" -  default for all patterns if not matched to any
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			// Or
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Ooops"))
			return
		}

		log.Printf("Data - %s\n", data)

		// Because responseWriter implements ioWriter
		fmt.Fprintf(w, "Hello %s\n", data)

		// Or using write method
		// w.Write([]byte(fmt.Sprintf("Hello %s", data)))
	})

	// HandleFunc converts a func to a handler and registers it to a ServeMux
	// (a server multiplexer -> basically mapping pattern/path to a func)
	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye World")
	})

	http.ListenAndServe(":9090", nil)

}
