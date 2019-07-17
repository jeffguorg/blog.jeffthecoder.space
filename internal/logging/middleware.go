package logging

import (
	"net/http"
	"time"
)

type ResponseWriter struct {
	writer http.ResponseWriter

	statusCode int
}

func (w *ResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *ResponseWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.statusCode = status
	w.writer.WriteHeader(status)
}

func (w *ResponseWriter) StatusCode() int {
	if w.statusCode == 0 {
		return 200
	}
	return w.statusCode
}

func WrapResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		writer: w,
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		ww := WrapResponseWriter(w)
		defer func() {
			Info("path: ", r.URL.Path, " | request time: ", time.Now().Sub(begin).String(), " | http status: ", ww.StatusCode())
		}()
		next.ServeHTTP(ww, r)
	})
}
