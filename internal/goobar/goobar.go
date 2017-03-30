package goobar

import "net/http"
import "log"
import "strings"
import "net/url"

type Action func(x Exchange) Responder

func RegisterAction(pattern string, action Action) {
	http.HandleFunc(pattern, makeHandler(action))
}

func RegisterFileServer(pattern string, rootPath string) {
	fileRootPath = rootPath
	handler := http.FileServer(http.Dir(rootPath))
	http.Handle(pattern, handler)
}

func RegisterViewPath(path string) {
	viewPath = path
}

var fileRootPath string
var viewPath string

func makeHandler(action Action) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer recoverFromActionPanic(w, r)
		x := makeExchange(w, r)
		responder := action(x)
		if ct := responder.ContentType(); len(strings.TrimSpace(ct)) > 0 {
			w.Header().Set("Content-Type", ct)
		}
		responder.Respond(w)
	}
}

func recoverFromActionPanic(w http.ResponseWriter, r *http.Request) {
	switch x := recover(); p := x.(type) {
	case nil: // nothing thrown, ignore
	case actionPanic:
		u, _ := url.QueryUnescape(r.URL.String())
		log.Printf("%s: %s", u, p.msg)
		http.NotFound(w, r)
	default:
		panic(x) // rethrow
	}
}
