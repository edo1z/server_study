package main

import (
	_ "./daytime"
	_ "./http"
	"./simple_chat"
)

func main() {
	//http.Server()
	//daytime.Server()
	simple_chat.Server()
}
