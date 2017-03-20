package goobar

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

type Responder interface {
	Respond(writer io.Writer) error
	ContentType() string
}

func JSON(v interface{}) Responder {
	return &jsonResponder{v}
}

type jsonResponder struct {
	value interface{}
}

func (r jsonResponder) Respond(writer io.Writer) error {
	//http.Handle("", http.FileServer(http.Dir("")))
	return json.NewEncoder(writer).Encode(r.value)
}

func (r jsonResponder) ContentType() string {
	return "application/json; charset=utf-8"
}

func PlainText(text string) Responder {
	return &plainTextResponder{text}
}

type plainTextResponder struct {
	text string
}

func (r plainTextResponder) Respond(writer io.Writer) error {
	_, err := io.WriteString(writer, r.text)
	return err
}

func (r plainTextResponder) ContentType() string {
	return "text/plain; charset=utf-8"
}

func XML(v interface{}) Responder {
	return &xmlResponder{v}
}

type xmlResponder struct {
	value interface{}
}

func (r xmlResponder) Respond(writer io.Writer) error {
	return xml.NewEncoder(writer).Encode(r.value)
}

func (r xmlResponder) ContentType() string {
	return "text/xml; charset=utf-8"
}

var nop = nopResponder{}

func Nop() Responder {
	return &nop
}

type nopResponder struct{}

func (r nopResponder) Respond(writer io.Writer) error {
	return nil
}

func (r nopResponder) ContentType() string {
	return ""
}

// func View(path string) Responder {
// }

// func ImagePNG(img image.Image) Responder {
// }
