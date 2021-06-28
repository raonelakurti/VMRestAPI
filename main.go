package main

import (
	"log"
	"fmt"
	"os"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	log.Printf("listening port: %s\n", "9000")

	//create data.json file if does not exists
	if _, err := os.Stat("data.json"); err == nil {
		fmt.Printf("File exists\n");  
	  } else {
		  //create a file with data
		  emptyFile, err := os.Create("data.json")
		  if err != nil {
			  log.Fatal(err)
		  }
		  emptyFile.Close()
	  
	  }

	//start router
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":9000", router))
}