package characters

import (
	"gtk/cmd/person"
	"gtk/cmd/weapon"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type Frag struct {
	P              *person.Person
	Life           int
	Damage         int
	MaxCountAttack int
	BlockAttack    bool
	CountAttack    int
}

func NewFrag(x, y, step, width, height int, path string, area *gtk.Fixed) *Hero {
	return &Hero{
		P:              person.NewPerson(x, y, step, width, height, path, area),
		Life:           100,
		Damage:         20,
		MaxCountAttack: 2,
		BlockAttack:    false,
		CountAttack:    0,
	}
}

func (h *Frag) TakeDamage(somePerson Hero, callback func()) {
	h.P.CheckCollision(*somePerson.P.Pointers, callback)
	h.Life -= somePerson.Damage
	if h.Life <= 0 {
		h.Death()
	}
}

func (h *Frag) Attack(x, y int, weapon weapon.Arrow) {
	if !h.BlockAttack {
		go func() {
			usrX := h.P.Pointers.X
			usrY := h.P.Pointers.Y
			h.CountAttack += 1
			if h.CountAttack >= h.MaxCountAttack {
				h.BlockAttack = true
			}
			for i := 0; i < x-usrX && i < y-usrY || i < weapon.AttackRange; i++ {
				usrX += 1 * h.P.Step
				usrY += 1 * h.P.Step
				dum, _ := gtk.ImageNewFromFile("../../../../../Downloads/fireball.png")
				h.P.Area.Put(dum, usrX, usrY)
				h.P.Area.ShowAll()
				time.Sleep(time.Millisecond * 100)
				dum.Destroy()
			}
			h.CountAttack -= 1
			h.BlockAttack = false
		}()
	}
}

func (f *Frag) Death() {
	f.P.Destroy()
}
