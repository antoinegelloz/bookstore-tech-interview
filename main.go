package main

import (
	router2 "bookstore/router"
	"fmt"
)

func main() {
	router := router2.NewRouter()
	if err := router.Run(); err != nil {
		fmt.Printf("api error: %s\n", err)
		return
	}
}
