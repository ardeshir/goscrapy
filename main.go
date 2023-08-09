package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// create the variables for the response & error
	var r *http.Response
	var err error

	var filename = "intro"
	// Request index.html from ardeshir.io
	r, err = http.Get("https://www.ardeshir.io/" + filename)

	// if there is a problem accessing the server, kill the program and print the error to console
	if err != nil {
		panic(err)
	}

	// Check the status code returned by the server
	if r.StatusCode == 200 {
		// The request was successful!
		var page []byte

		// we know the size of the respon is 1270
		var bodyLen int = 16320 // 9180 // #16320

		page = make([]byte, bodyLen)

		// read the data from server
		r.Body.Read(page)

		// Open a writable file on your computer (create if it does not exits)
		var out *os.File

		out, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		// write the content
		out.Write(page)
		out.Close()

	} else {
		log.Fatal("Failed to retrieve the web page. Recieved status code ", r.Status)
	}

} // end of main
