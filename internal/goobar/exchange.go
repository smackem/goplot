package goobar

import (
	"fmt"
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

func (x *Exchange) GetId() (int, bool) {
	id, err := strconv.Atoi(x.id)
	return id, err == nil
}

func (x *Exchange) MustGetId() int {
	id, err := strconv.Atoi(x.id)
	if err != nil {
		x.doPanic("id not present or in the wrong format")
	}
	return id
}

func (x *Exchange) GetIdString() string {
	return x.id
}

func (x *Exchange) MustGetIdString() string {
	if x.id == "" {
		panic("id not present")
	}
	return x.id
}

func (x *Exchange) GetInt(key string) (int, bool) {
	s := x.r.FormValue(key)
	if s == "" {
		return 0, false
	}
	val, err := strconv.Atoi(s)
	return val, err == nil
}

func (x *Exchange) MustGetInt(key string) int {
	s := x.r.FormValue(key)
	val, err := strconv.Atoi(s)
	if err != nil {
		x.doPanic(fmt.Sprintf("Value %s not present or not an integer", key))
	}
	return val
}

func (x *Exchange) GetString(key string) string {
	return x.r.FormValue(key)
}

func (x *Exchange) MustGetString(key string) string {
	s := x.r.FormValue(key)
	if s == "" {
		x.doPanic(fmt.Sprintf("Value %s not present", key))
	}
	return s
}

func (x *Exchange) GetFloat(key string) (float64, bool) {
	s := x.r.FormValue(key)
	val, err := strconv.ParseFloat(s, 64)
	return val, err == nil
}

func (x *Exchange) MustGetFloat(key string) float64 {
	s := x.r.FormValue(key)
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		x.doPanic(fmt.Sprintf("Value %s not present or not a float", key))
	}
	return val
}

func (x *Exchange) GetBool(key string) (bool, bool) {
	s := x.r.FormValue(key)
	val, err := strconv.ParseBool(s)
	return val, err == nil
}

func (x *Exchange) MustGetBool(key string) bool {
	s := x.r.FormValue(key)
	val, err := strconv.ParseBool(s)
	if err != nil {
		x.doPanic(fmt.Sprintf("Value %s not present or not a boolean", key))
	}
	return val
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
