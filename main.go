package main

import (
	"bookstore/router"
	"fmt"
)

func main() {
	r := router.NewRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("api error: %s\n", err)
		return
	}
}
