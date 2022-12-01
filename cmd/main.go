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
	log.SetFlags(0)

	address := flag.String("a", ":8000", "Listen address")
	flag.Parse()

	server, err := n2i.NewServer(&n2i.Config{
		LogHandler: slog.NewTextHandler(os.Stderr),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listen", *address)
	log.Fatal(http.ListenAndServe(*address, server))
}
