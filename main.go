package main

import (
	"./daytime"
	_ "./http"
)

func main() {
	//http.Server()
	daytime.Server()
}
