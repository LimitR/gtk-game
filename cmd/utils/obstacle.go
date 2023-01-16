package utils

import (
	"log"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

func FixedImageByName(b *gtk.Builder, name string, count int, texturePath string) {
	for i := 1; i <= count; i++ {
		nameObject := name + "-" + strconv.Itoa(i)
		obj, err := b.GetObject(nameObject)
		if err != nil {
			log.Fatal("Image")
		}
		img := obj.(*gtk.Image)
		img.SetFromFile(texturePath)
	}
}
