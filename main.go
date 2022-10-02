package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//Setting Host address & Port information.
	port := ":8080"

	//This request can come to any interface, hence setting 0.0.0.0
	address := "0.0.0.0" + port
	http.HandleFunc("/", handler)

	//Starting web server here, if there is an error, going to be dealt
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

//the main logic here, returning timestamp and hostname of the current environment
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(getTimeStamp() + "\n" + getHostName() + "\n"))
}

//this function is to get current timestamp from environment.
func getTimeStamp() string {
	return time.Now().Format(time.UnixDate)
}

//this function is to get hostname of current environment. It would return error if hostname can not be retrieved.
func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return hostname
}
