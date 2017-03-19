package goobar

import "io"
import "net/http"
import "log"

type Action func(x Exchange) Encoder

func RegisterAction(pattern string, action Action) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			x := recover()
			switch p := x.(type) {
			case nil: // nothing thrown, ignore
			case actionPanic:
				log.Println(p.msg)
				http.NotFound(w, r)
			default:
				panic(x) // rethrow
			}
		}()
		x := makeExchange(w, r)
		Encoder := action(x)
		Encoder.Encode(w)
	})
}
