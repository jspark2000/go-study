package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jspark2000/go-study/src/router"
)

func main() {
	router := router.NewBasicJudgeRouter()
	http.HandleFunc("/judge", router.HandleJudge)

	fmt.Println("Server is Running at 4001...")

	if err := http.ListenAndServe(":4001", nil); err != nil {
		log.Fatal(err)
	}
}
