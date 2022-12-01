package main

import (
	"flag"
	"fmt"
	"github.com/HuguesGuilleus/n2i_2022"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"os"
)

func main() {
	address := flag.String("a", ":8000", "Listen address")
	flag.Parse()

	config := n2i.Config{
		LogHandler: slog.NewTextHandler(os.Stdout),
	}

	fmt.Println("Listen", *address)
	log.Fatal(http.ListenAndServe(*address, n2i.NewServer(&config)))
}
