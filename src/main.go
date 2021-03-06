package main

import (
	"fmt"
	"flag"
	"log"
	"net/http"
	"github.com/stevenandrewcarter/gouter"
	"github.com/stevenandrewcarter/gouter/controllers"
	"github.com/stevenandrewcarter/gouter/lib"
)

// Loads the parameters that got provided on the command line. If not provided will use the defaults instead
func loadParameters() (string) {
	wordPtr := flag.String("port", gouter.Configuration().Application.Port, "Port Number for the Gouter to run at")
	flag.Parse()
	return *wordPtr
}

// Displays the start blurb for the router
func start(port string) {
	log.Printf("\n  ________               __                \n /  _____/  ____  __ ___/  |_  ___________ \n/   \\  ___ /  _ \\|  |  \\   __\\/ __ \\_  __ \\ \n\\    \\_\\  (  <_> )  |  /|  | \\  ___/|  | \\/\n \\______  /\\____/|____/ |__|  \\___  >__|\n        \\/                        \\/")
	log.Printf("Starting Gouter v0.5. A simple HTTP router for RESTful API calls.")
	log.Printf("Please call http://localhost:%v%v to configure Gouter.", port, gouter.Configuration().Application.AdminUrl)
	http.HandleFunc("/", lib.HandleRequest)
	controllers.Load()
	log.Printf("Listening for HTTP requests on Port '%v'", port)
}

// Main entry point for the Gouter project. Will listen on the configured port and models
// the http request to the matched request. The response will be returned to the original
// caller. 
func main() {
	port := loadParameters()
	start(port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
