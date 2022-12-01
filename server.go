package n2i

import (
	"github.com/HuguesGuilleus/n2i_2022/front"
	"net/http"
	"strconv"
)

type Config struct {
}

func NewServer(config *Config) http.Handler {
	return &server{}
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		serveStatic(w, "text/html", front.INDEX)
	case "/css":
		serveStatic(w, "text/css", front.CSS)
	case "/js":
		serveStatic(w, "application/javascript", front.JS)
	default:
		http.NotFound(w, r)
	}
}

func serveStatic(w http.ResponseWriter, contentType string, payload []byte) {
	w.Header().Add("Content-Type", contentType)
	w.Header().Add("Content-Length", strconv.Itoa(len(payload)))
	w.Write(payload)
}
