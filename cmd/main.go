package main

import (
	// "pkg/golang.org/x/exp/slog"
	"flag"
	"fmt"
	"github.com/HuguesGuilleus/n2i_2022"
	"log"
	"net/http"
)

func main() {
	address := flag.String("a", ":8000", "Listen address")
	flag.Parse()

	config := n2i.Config{}

	fmt.Println("Listen", *address)
	log.Fatal(http.ListenAndServe(*address, n2i.NewServer(&config)))
}
