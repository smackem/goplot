package goobar

import "net/http"
import "log"
import "strings"
import "net/url"
import "path/filepath"

type Handler struct {
	Get     Action
	Post    Action
	Put     Action
	Delete  Action
	Options Action
}

type Action func(x Exchange) Responder

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer recoverFromActionPanic(w, r)
	x := makeExchange(w, r)
	var responder Responder

	switch r.Method {
	case "GET":
		if h.Get == nil {
			http.NotFound(w, r)
			return
		}
		responder = h.Get(x)
	case "POST":
		if h.Post == nil {
			http.NotFound(w, r)
			return
		}
		responder = h.Post(x)
	case "PUT":
		if h.Put == nil {
			http.NotFound(w, r)
			return
		}
		responder = h.Put(x)
	case "DELETE":
		if h.Delete == nil {
			http.NotFound(w, r)
			return
		}
		responder = h.Delete(x)
	case "OPTIONS":
		if h.Options == nil {
			http.NotFound(w, r)
			return
		}
		responder = h.Options(x)
	default:
		http.NotFound(w, r)
		return
	}

	if ct := responder.ContentType(); len(strings.TrimSpace(ct)) > 0 {
		w.Header().Set("Content-Type", ct)
	}
	responder.Respond(w)
}

func Get(action Action) *Handler {
	return &Handler{Get: action}
}

func Post(action Action) *Handler {
	return &Handler{Post: action}
}

func Put(action Action) *Handler {
	return &Handler{Put: action}
}

func Delete(action Action) *Handler {
	return &Handler{Delete: action}
}

func AnyMethod(action Action) *Handler {
	return &Handler{
		Get:     action,
		Post:    action,
		Put:     action,
		Delete:  action,
		Options: action,
	}
}

func SetViewFolder(folder string) {
	viewFolder = filepath.Clean(folder)
}

func GetViewFolder() string {
	return viewFolder
}

var viewFolder = "resource/view"

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
