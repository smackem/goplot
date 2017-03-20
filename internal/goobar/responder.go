package goobar

import (
	"encoding/json"
	"encoding/xml"
	"image"
	"image/png"
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

// func View(path string) Responder {
// }
