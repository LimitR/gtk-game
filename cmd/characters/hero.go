package characters

import (
	"gtk/cmd/person"
	"gtk/cmd/weapon"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

type Hero struct {
	P              *person.Person
	Life           int
	Damage         int
	MaxCountAttack int
	BlockAttack    bool
	CountAttack    int
}

func NewHero(x, y, step, width, height int, path string, area *gtk.Fixed) *Hero {
	return &Hero{
		P:              person.NewPerson(x, y, step, width, height, path, area),
		Life:           100,
		Damage:         5,
		MaxCountAttack: 2,
		BlockAttack:    false,
		CountAttack:    0,
	}
}

func (h *Hero) TakeDamage(somePerson Frag, callback func()) {
	h.P.CheckCollision(*somePerson.P.Pointers, func() {
		h.Life -= somePerson.Damage
		callback()
	})
}

func (h *Hero) Attack(x, y int, weapon weapon.Arrow) {
	if !h.BlockAttack {
		go func() {
			defer h.endAttack()
			defer weapon.Destroy()
			weapon.X = h.P.Pointers.X
			weapon.Y = h.P.Pointers.Y
			h.CountAttack += 1
			if h.CountAttack >= h.MaxCountAttack {
				h.BlockAttack = true
			}
			if x-weapon.X > 50 {
				if y-weapon.Y > 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.X += 1 * 10
						weapon.Y += 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return
				}
				if weapon.Y-y > 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.X += 1 * 10
						weapon.Y -= 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return
				}
				if y == weapon.Y || y-weapon.Y < 50 || y-weapon.Y > 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.X += 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return
				}
				return
			}
			if weapon.X-x > 50 {
				if y-weapon.Y > 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.X -= 1 * 10
						weapon.Y += 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return
				}
				if weapon.Y-y > 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.X -= 1 * 10
						weapon.Y -= 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return
				}
				if y == weapon.Y || weapon.Y-y < 50 || y-weapon.Y < 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.X -= 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return
				}
				return
			}
			if x == weapon.X || x-weapon.X < 50 || weapon.X-x < 50 {
				if y-weapon.Y > 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.Y += 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return

				}
				if weapon.Y-y > 50 {
					for i := 0; i < weapon.AttackRange; i++ {
						weapon.Y -= 1 * 10
						dum := weapon.GetTexture()
						h.P.Area.Put(dum, weapon.X, weapon.Y)
						dum.Show()
						time.Sleep(time.Millisecond * 100)
						dum.Destroy()
					}
					return
				}
				return
			}
		}()
	}
}

func (h *Hero) endAttack() {
	h.CountAttack -= 1
	h.BlockAttack = false
}
