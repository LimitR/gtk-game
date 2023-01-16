package utils

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func ExitIfError[T gtk.Widget](g T, e error) T {
	if e != nil {
		log.Fatalln(e)
	}
	return g
}
