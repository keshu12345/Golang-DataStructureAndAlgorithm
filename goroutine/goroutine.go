package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	_ "runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})
	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
		}
	}()
	for {
		runtime.Gosched()
	}
}
