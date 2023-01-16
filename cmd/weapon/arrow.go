package weapon

import (
	"gtk/cmd/utils"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

type Arrow struct {
	AttackRange       int
	Damage            int
	DamageTexturePath string
	Texture           *gtk.Image
	path              string
	X, Y              int
}

func NewArraw(path string) *Arrow {
	t, err := gtk.ImageNewFromFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return &Arrow{
		AttackRange:       10,
		Damage:            10,
		DamageTexturePath: path,
		Texture:           t,
		path:              path,
		X:                 0,
		Y:                 0,
	}
}

func (a *Arrow) GetTexture() *gtk.Image {
	t, err := gtk.ImageNewFromFile(a.path)
	if err != nil {
		log.Fatalln(err)
	}
	a.Texture = t
	return a.Texture

}

func (p *Arrow) CheckCollision(somePerson utils.Pointers, callback func()) {
	if p.X >= somePerson.X && p.X <= somePerson.X+somePerson.Width {
		if p.Y >= somePerson.Y && p.Y <= somePerson.Y+somePerson.Height {
			callback()
			return
		}
	}
	if p.Y >= somePerson.Y && p.Y <= somePerson.Y+somePerson.Height {
		if p.X >= somePerson.X && p.X <= somePerson.X+somePerson.Width {
			callback()
			return
		}
	}
}

func (a *Arrow) Destroy() {
	a.Texture.Destroy()
}
