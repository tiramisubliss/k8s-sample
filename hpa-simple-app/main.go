package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
)

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var x float64 = 0.0001
	for i := 0; i < 1000000; i++ {
		x += sqrt(x)
	}
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hola World, Here from Kubernetes!\n")
	fmt.Fprintf(w, "Version: 2.0.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
