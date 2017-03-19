package goobar

import "io"

type Result interface {
	Encode(writer io.Writer) error
}

type Action func(x Exchange) Result

func RegisterAction(pattern string, action Action) {
}
