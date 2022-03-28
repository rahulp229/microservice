package main

import (
	"log"
	"microservice/router"
	"net/http"
)

func main() {
	routes := router.SetRoutes()

	log.Fatal(http.ListenAndServe(":9090", routes))
}

/*
Problem
url- https://min-api.cryptocompare.com/data/pricemultifull?fsyms=BTC&tsyms=USD,EUR

*/
