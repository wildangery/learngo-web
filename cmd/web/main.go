package main

import (
	"fmt"
	"net/http"

	"github.com/wildangery/learngo-web/packages/handlers"
)

var portNumber = ":3000"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/abouts", handlers.About)
	fmt.Println(fmt.Sprintf("Starting application osss port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
