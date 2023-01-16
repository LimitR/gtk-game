package elements

import (
	"gtk/cmd/utils"

	"github.com/gotk3/gotk3/gtk"
)

type IWall interface {
}

type Wall struct {
	Pointers *utils.Pointers
	path     string
	object   *gtk.Image
	area     *gtk.Fixed
}

func NewWall(x, y, width, height int, path string, area *gtk.Fixed) *Wall {
	object, _ := gtk.ImageNewFromFile(path)
	return &Wall{&utils.Pointers{x, y, width, height}, path, object, area}
}

func (p *Wall) Render() {
	p.object.Destroy()
	p.object, _ = gtk.ImageNewFromFile(p.path)
	p.area.Put(p.object, p.Pointers.X, p.Pointers.Y)
	p.area.ShowAll()
}
