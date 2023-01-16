package person

import (
	"gtk/cmd/utils"

	"github.com/gotk3/gotk3/gtk"
)

type IPerson interface {
	MovePosition(x, y int)
	StepUp()
	StepDown()
	StepLeft()
	StepRight()
	Render()
	ChangeRender(path string)
	CheckCollision(somePerson utils.Pointers, callback func())
	Destroy()
}

type Person struct {
	Pointers *utils.Pointers
	size     int
	Step     int
	path     string
	object   *gtk.Image
	Area     *gtk.Fixed
}

func NewPerson(x, y, step, width, height int, path string, area *gtk.Fixed) *Person {
	person, _ := gtk.ImageNewFromFile(path)
	return &Person{
		Step:   step,
		path:   path,
		object: person,
		Area:   area,
		size:   width * height,
		Pointers: &utils.Pointers{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
	}
}

func (p *Person) MovePosition(x, y int) {
	p.Pointers.X = x
	p.Pointers.Y = y
	p.Render()
}

func (p *Person) StepUp() {
	p.Pointers.Y -= p.Step
	p.Render()
}

func (p *Person) StepDown() {
	p.Pointers.Y += p.Step
	p.Render()
}

func (p *Person) StepLeft() {
	p.Pointers.X -= p.Step
	p.Render()
}

func (p *Person) StepRight() {
	p.Pointers.X += p.Step
	p.Render()
}

func (p *Person) Render() {
	p.object.Destroy()
	p.object, _ = gtk.ImageNewFromFile(p.path)
	p.Area.Put(p.object, p.Pointers.X, p.Pointers.Y)
	p.object.Show()
}

func (p *Person) ChangeRender(path string) {
	p.object.Destroy()
	p.object, _ = gtk.ImageNewFromFile(path)
	p.Area.Put(p.object, p.Pointers.X, p.Pointers.Y)
	p.object.Show()
}

func (p *Person) CheckCollision(somePerson utils.Pointers, callback func()) {
	if p.Pointers.X >= somePerson.X && p.Pointers.X <= somePerson.X+somePerson.Width {
		if p.Pointers.Y >= somePerson.Y && p.Pointers.Y <= somePerson.Y+somePerson.Height {
			callback()
			return
		}
	}
	if p.Pointers.Y >= somePerson.Y && p.Pointers.Y <= somePerson.Y+somePerson.Height {
		if p.Pointers.X >= somePerson.X && p.Pointers.X <= somePerson.X+somePerson.Width {
			callback()
			return
		}
	}
}

func (p *Person) Destroy() {
	p.object.Destroy()
}
