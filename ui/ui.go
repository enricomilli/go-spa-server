package ui

import (
	"embed"
	"io/fs"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
)

//go:embed dist
var distFS embed.FS

func SetupRoutes(router *chi.Mux) {
    // Strip the "dist" prefix to serve files from root
    fsys, err := fs.Sub(distFS, "dist")
    if err != nil {
        panic(err)
    }

    fileServer := http.FileServer(http.FS(fsys))
    router.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)

        if !fileMatcher.MatchString(r.URL.Path) {
            // Serve index.html for non-file paths
            content, err := distFS.ReadFile("dist/index.html")
            if err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                return
            }
            w.Header().Set("Content-Type", "text/html; charset=utf-8")
            w.Write(content)
            return
        }

        fileServer.ServeHTTP(w, r)
    }))
}
