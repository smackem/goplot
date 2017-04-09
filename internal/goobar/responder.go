package goobar

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
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
	return json.NewEncoder(writer).Encode(r.value)
}

func (r jsonResponder) ContentType() string {
	return "application/json; charset=utf-8"
}

func PlainText(text string) Responder {
	return plainTextResponder(text)
}

type plainTextResponder string

func (r plainTextResponder) Respond(writer io.Writer) error {
	_, err := io.WriteString(writer, string(r))
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

func ImagePNG(img image.Image) Responder {
	return &pngResponder{img}
}

type pngResponder struct {
	img image.Image
}

func (r pngResponder) Respond(writer io.Writer) error {
	return png.Encode(writer, r.img)
}

func (r pngResponder) ContentType() string {
	return "image/png"
}

func Binary(reader io.Reader) Responder {
	return &binaryResponder{reader}
}

type binaryResponder struct {
	reader io.Reader
}

func (r binaryResponder) Respond(writer io.Writer) error {
	_, err := io.Copy(writer, r.reader)
	return err
}

func (r binaryResponder) ContentType() string {
	return "application/octet-stream"
}

func Error(statusCode int, message string) Responder {
	return &errorResponder{statusCode, message}
}

type errorResponder struct {
	statusCode int
	message    string
}

func (r errorResponder) Respond(writer io.Writer) error {
	if w, ok := writer.(http.ResponseWriter); ok {
		http.Error(w, r.message, r.statusCode)
		return nil
	}
	_, err := fmt.Fprintf(writer, "Error %d: %s", r.statusCode, r.message)
	return err
}

func (r errorResponder) ContentType() string {
	return "text/plain; charset=utf-8"
}

func View(path string, model interface{}) Responder {
	return &viewResponder{path, model}
}

type viewResponder struct {
	path  string
	model interface{}
}

func (r viewResponder) Respond(writer io.Writer) error {
	templ, name, err := getTemplate(r.path)
	if err != nil {
		log.Print(err)
		return err
	}
	return templ.ExecuteTemplate(writer, name, r.model)
}

func (r viewResponder) ContentType() string {
	return "text/html; charset=utf-8"
}
