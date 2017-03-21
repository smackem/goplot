package goobar

import "net/http"
import "log"
import "strings"

type Action func(x Exchange) Responder

func RegisterAction(pattern string, action Action) {
	http.HandleFunc(pattern, makeHandler(action))
}

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
		log.Println(p.msg)
		http.NotFound(w, r)
	default:
		panic(x) // rethrow
	}
}
