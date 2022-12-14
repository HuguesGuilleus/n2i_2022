package n2i

import (
	"bytes"
	"compress/gzip"
	"embed"
	"fmt"
	"github.com/HuguesGuilleus/n2i_2022/front"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

type Config struct {
	// The directory to store the database.
	DBDirectory string
	// Handler to store the logger.
	LogHandler slog.Handler
}

//go:embed jeu
var jeu embed.FS

func (s *server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Print("DEBUG ")
	fmt.Println(r.URL.Path)

	if r.URL.Path == "/page/" || r.URL.Path == "/page" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
	}

	// send all data of this page
	var id int
	_, err := fmt.Sscanf(r.URL.Path, "/page/%d", &id)
	if err != nil {
		http.Error(w, "Need a good id", 400)
		return
	}

	page, err := s.db.loadPage(id)
	if err != nil {
		http.Error(w, "Page not found", 404)
		return
	}

	front.PageTemplate.Execute(w, page)
}

func NewServer(config *Config) (http.Handler, error) {
	db, err := newDB(config.DBDirectory)
	if err != nil {
		return nil, err
	}

	_, titles, _ := db.loadAllMetadata()

	s := &server{
		log: slog.New(config.LogHandler),
		db:  db,
	}

	s.mux.Handle("/css", staticHandler("text/css", front.CSS))
	s.mux.Handle("/js", staticHandler("application/javascript", front.JS))
	s.mux.Handle("/img/", http.FileServer(http.FS(front.Images)))

	s.mux.Handle("/jeu/", http.FileServer(http.FS(jeu)))

	s.mux.HandleFunc("/page/", s.HandleRequest)
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		front.IndexTemplate.Execute(w, titles)
	})

	return s, nil
}

type server struct {
	log *slog.Logger
	db  database
	mux http.ServeMux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.log.Info("http", "url", r.URL.String())
	s.mux.ServeHTTP(w, r)
}

func staticHandler(contentType string, payload []byte) http.Handler {
	gzipBuffer := bytes.Buffer{}
	gzipWriter, _ := gzip.NewWriterLevel(&gzipBuffer, gzip.BestCompression)
	gzipWriter.Write(payload)
	gzipWriter.Close()
	gzipBytes := gzipBuffer.Bytes()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", contentType)

		acceptGzip := false
		for _, acceptEncoding := range r.Header["Accept-Encoding"] {
			for _, encoding := range strings.FieldsFunc(acceptEncoding, splitSpaceComa) {
				acceptGzip = acceptGzip || encoding == "gzip"
			}
		}

		if acceptGzip {
			w.Header().Add("Content-Encoding", "gzip")
			w.Write(gzipBytes)
		} else {
			w.Header().Add("Content-Length", strconv.Itoa(len(payload)))
			w.Write(payload)
		}
	})
}

func splitSpaceComa(r rune) bool { return r == ',' || unicode.IsSpace(r) }
