// main.go

package main

import (
	"fmt"
	"net/http"
	"shopping-app/api" // Importing your custom package
	"time"
)

func main() {

	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
