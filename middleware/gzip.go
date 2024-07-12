package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

var gzipWriterPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(io.Discard)
	},
}

type gzipResponseWriter struct {
	http.ResponseWriter
	gzWriter *gzip.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	return w.gzWriter.Write(b)
}

func (w *gzipResponseWriter) Close() error {
	return w.gzWriter.Close()
}

// GzipMiddleware compresses HTTP responses with gzip
func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Set headers before gzipping the response
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Del("Content-Length")

		gz := gzipWriterPool.Get().(*gzip.Writer)
		defer gzipWriterPool.Put(gz)
		gz.Reset(w)
		defer gz.Close()

		gzw := &gzipResponseWriter{ResponseWriter: w, gzWriter: gz}

		next.ServeHTTP(gzw, r)
	})
}

// CompressStaticFileHandler compresses static files if the client supports gzip
func CompressStaticFileHandler(next http.Handler) http.Handler {
	return GzipMiddleware(next)
}
