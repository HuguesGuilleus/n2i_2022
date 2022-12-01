package n2i

import (
	"bytes"
	"compress/gzip"
	"github.com/HuguesGuilleus/n2i_2022/front"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"embed"
)

type Config struct {
	// The directory to store the database.
	DBDirectory string
	// Handler to store the logger.
	LogHandler slog.Handler
}

//go:embed jeu
var jeu embed.FS

func NewServer(config *Config) (http.Handler, error) {
	db, err := newDB(config.DBDirectory)
	if err != nil {
		return nil, err
	}

	s := &server{
		log: slog.New(config.LogHandler),
		db:  db,
	}

	s.mux.Handle("/", staticHandler("text/html", front.INDEX))
	s.mux.Handle("/css", staticHandler("text/css", front.CSS))
	s.mux.Handle("/js", staticHandler("application/javascript", front.JS))

	s.mux.Handle("/jeu/", http.FileServer(http.FS(jeu)))

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
