package n2i

import (
	"github.com/HuguesGuilleus/n2i_2022/front"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

type Config struct {
	LogHandler slog.Handler
}

func NewServer(config *Config) http.Handler {
	return &server{
		log: slog.New(config.LogHandler),
	}
}

type server struct {
	log *slog.Logger
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.log.Info("http", "url", r.URL.String())

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
