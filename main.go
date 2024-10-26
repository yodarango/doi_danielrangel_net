package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type DOI map[string]string

func main (){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		doiPath := r.URL.Path

		if doiPath == "/" {
			http.Error(w, "This link does not exists. Please verify your URL", 404)
			return 
		}

		// Open the json file
		file, err := os.Open("./doi.json")

		if err !=nil {
			log.Fatalf("I could not open the file %v", err)
		}

		defer file.Close()

		// Read the file
		data, err := io.ReadAll(file)

		if err != nil {
			log.Fatalf("I could not read the contents of the file %v", err)
		}

		// Parse the file
		var doi DOI

		err = json.Unmarshal(data, &doi) 
		
		if err != nil {
			log.Fatalf("Error unmarshaling file %v", err)
		}

		cleanDoiPath := strings.TrimPrefix(doiPath, "/")
		redirectUrl := doi[cleanDoiPath]

		if redirectUrl == "" {
			http.Error(w, "This link does not exists. Please verify your URL", 404)
			return 
		}

		http.Redirect(w, r, redirectUrl, 303)
	})

	http.ListenAndServe(":8004", nil)
}