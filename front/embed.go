package front

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"io/fs"
)

var (
	//go:embed img
	Images embed.FS

	//go:embed js
	jsFS embed.FS
	//go:embed css
	cssFS embed.FS

	JS  []byte = merge(jsFS, js.Minify)
	CSS []byte = merge(cssFS, css.Minify)

	//go:embed index.html
	INDEX []byte
)

func merge(fsys fs.FS, minifyFunc minify.MinifierFunc) []byte {
	source := bytes.Buffer{}
	err := fs.WalkDir(fsys, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Walk in %q: %w", path, err)
		} else if entry.IsDir() {
			return nil
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return fmt.Errorf("Read %q: %w", path, err)
		}
		source.Write(data)
		source.WriteByte('\n')

		return nil
	})
	if err != nil {
		panic(err)
	}

	output := bytes.Buffer{}
	if err := minifyFunc(nil, &output, &source, nil); err != nil {
		panic(err)
	}
	return output.Bytes()
}

// Minify HTML
func init() {
	buff := bytes.Buffer{}
	if err := html.Minify(nil, &buff, bytes.NewReader([]byte(INDEX)), nil); err != nil {
		panic(err)
	}
	INDEX = buff.Bytes()
}
