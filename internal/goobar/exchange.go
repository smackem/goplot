package goobar

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Exchange struct {
	w  http.ResponseWriter
	r  *http.Request
	id string
}

func (x Exchange) Request() *http.Request {
	return x.r
}

func (x *Exchange) GetID() (int, bool) {
	log.Printf("ID: %s -> int", x.id)
	id, err := strconv.Atoi(x.id)
	return id, err == nil
}

func (x *Exchange) MustGetID() int {
	log.Printf("ID: %s -> int", x.id)
	id, err := strconv.Atoi(x.id)
	if err != nil {
		x.doPanic("id not present or in the wrong format")
	}
	return id
}

func (x *Exchange) GetIDString() string {
	log.Printf("ID: %s -> string", x.id)
	return x.id
}

func (x *Exchange) MustGetIDString() string {
	log.Printf("ID: %s -> string", x.id)
	if x.id == "" {
		panic("id not present")
	}
	return x.id
}

func (x *Exchange) GetInt(key string) (int, bool) {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> int", key, s)
	if s == "" {
		return 0, false
	}
	val, err := strconv.Atoi(s)
	return val, err == nil
}

func (x *Exchange) MustGetInt(key string) int {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> int", key, s)
	val, err := strconv.Atoi(s)
	if err != nil {
		x.doPanic(fmt.Sprintf("value %s not present or not an integer", key))
	}
	return val
}

func (x *Exchange) GetIntOrDefault(key string, defaultVal int) int {
	if val, ok := x.GetInt(key); ok {
		return val
	}
	return defaultVal
}

func (x *Exchange) GetString(key string) string {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> string", key, s)
	return s
}

func (x *Exchange) MustGetString(key string) string {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> string", key, s)
	if s == "" {
		x.doPanic(fmt.Sprintf("value %s not present", key))
	}
	return s
}

func (x *Exchange) GetStringOrDefault(key string, defaultVal string) string {
	if val := x.GetString(key); val != "" {
		return val
	}
	return defaultVal
}

func (x *Exchange) GetFloat(key string) (float64, bool) {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> float", key, s)
	val, err := strconv.ParseFloat(s, 64)
	return val, err == nil
}

func (x *Exchange) MustGetFloat(key string) float64 {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> float", key, s)
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		x.doPanic(fmt.Sprintf("value %s not present or not a float", key))
	}
	return val
}

func (x *Exchange) GetFloatOrDefault(key string, defaultVal float64) float64 {
	if val, ok := x.GetFloat(key); ok {
		return val
	}
	return defaultVal
}

func (x *Exchange) GetBool(key string) (bool, bool) {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> bool", key, s)
	val, err := strconv.ParseBool(s)
	return val, err == nil
}

func (x *Exchange) MustGetBool(key string) bool {
	s := x.r.FormValue(key)
	log.Printf("%s: %s -> bool", key, s)
	val, err := strconv.ParseBool(s)
	if err != nil {
		x.doPanic(fmt.Sprintf("value %s not present or not a boolean", key))
	}
	return val
}

func (x *Exchange) GetBoolOrDefault(key string, defaultVal bool) bool {
	if val, ok := x.GetBool(key); ok {
		return val
	}
	return defaultVal
}

///////////////////////////////////////////////////////////////////////////////

type actionPanic struct {
	msg string
}

func (x *Exchange) doPanic(msg string) {
	panic(actionPanic{msg})
}

func makeExchange(w http.ResponseWriter, r *http.Request) Exchange {
	_, id := path.Split(r.URL.Path)

	return Exchange{
		w:  w,
		r:  r,
		id: id,
	}
}
