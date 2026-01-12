package utility

import (
	"bufio"
	"io"
	"net"
	"net/http"
)

type RPWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *RPWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *RPWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *RPWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *RPWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
