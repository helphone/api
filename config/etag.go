package config

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ETag() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer = &etagWriter{c.Writer, c.Request}
		c.Next()
	}
}

type etagWriter struct {
	gin.ResponseWriter
	req *http.Request
}

func (e *etagWriter) Write(data []byte) (int, error) {
	if e.req.Method == "GET" && len(data) > 1024 {
		hash := md5.New()
		for header := range e.Header() {
			io.WriteString(hash, header)
		}
		io.WriteString(hash, string(data))
		etag := fmt.Sprintf("%x", hash.Sum(nil))

		if match, ok := e.req.Header["If-None-Match"]; ok {
			if match[0] == etag {
				e.WriteHeader(http.StatusNotModified)
				data = []byte{}
			}
		} else {
			e.Header().Set("ETag", etag)
		}
	}

	size, err := e.ResponseWriter.Write(data)
	return size, err
}
