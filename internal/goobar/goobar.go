package goobar

import "net/http"
import "log"
import "strings"
import "net/url"
import "path/filepath"

type Handler struct {
	Get, Post, Put, Delete, Options Action
}

type Action func(x Exchange) Responder

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer recoverFromActionPanic(w, r)

	action := h.getAction(r.Method)
	if action == nil {
		http.NotFound(w, r)
		return
	}

	x := makeExchange(w, r)
	responder := action(x)
	if ct := responder.ContentType(); len(strings.TrimSpace(ct)) > 0 {
		w.Header().Set("Content-Type", ct)
	}
	responder.Respond(w)
}

func (h *Handler) getAction(method string) Action {
	switch method {
	case "GET":
		return h.Get
	case "POST":
		return h.Post
	case "PUT":
		return h.Put
	case "DELETE":
		return h.Delete
	case "OPTIONS":
		return h.Options
	default:
		return nil
	}
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

func ViewFolder() string {
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
