package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
)

type Kubus struct {
	Sisi float64
}

func (k *Kubus) Volume() (float64, error) {
	if k.Sisi == 4 {
		return 0, errors.New("angka sial")
	}
	return math.Pow(k.Sisi, 3), nil
}

func (k *Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

func (k *Kubus) Keliling() float64 {
	return k.Sisi * 12
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World!")
		fmt.Fprintln(w, "Ok Google !")
	})

	http.ListenAndServe(":"+getPort(), nil)
}
