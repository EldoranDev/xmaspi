package api

import (
	"github.com/EldoranDev/xmaspi/v2/static"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

var _ http.Handler = (*spaHandler)(nil)

func NewSpaHandler() http.Handler {
	return &spaHandler{}
}

type spaHandler struct {
}

func (s *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(".", path)

	_, err = fs.Stat(static.Files, path)
	if os.IsNotExist(err) {
		_, _ = w.Write([]byte(static.Index))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.FileServer(http.FS(static.Files)).ServeHTTP(w, r)
}
