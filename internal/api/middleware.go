package api

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jhachmer/imgo/internal/config"
	"github.com/jhachmer/imgo/internal/utils"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Logging logs all requests with its method, path and the time it took to process
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.Method, r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

// ValidateContentType validates content type of request
// Continues if request is of specified type
// Sends error as JSON if unsupported type
func ValidateContentType(ct string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			key := http.CanonicalHeaderKey("Content-Type")
			if r.Header.Get(key) != ct {
				errMsg := utils.ErrorResponse{Error: "Unsupported Media Type"}
				log.Println(errMsg.ErrorMessage())
				utils.WriteJSON(w, http.StatusUnsupportedMediaType, errMsg)
				return
			}
			f(w, r)
		}
	}
}

func clearTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// don't mess with file server urls
		if strings.HasPrefix(r.URL.Path, "/files") {
			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path != config.V1PREFIX+"/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func logFileServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/files") {
			log.Println(r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// TODO: implement
func Authenticate() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(w, r)
		}
	}
}
