package main

import (
	"fmt"
	"gtk/cmd/characters"
	"gtk/cmd/weapon"
	"log"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	b, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	err = b.AddFromFile("./main.glade")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	obj, err := b.GetObject("main")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}

	win := obj.(*gtk.Window)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	obj, _ = b.GetObject("fix")
	game := obj.(*gtk.Fixed)

	obj, err = b.GetObject("gl-area-1")
	if err != nil {
		log.Fatal("Ошибка:", err)
	}
	area := obj.(*gtk.GLArea)

	area.Connect("button-press-event", func(da *gtk.GLArea, cr *cairo.Context) {
		cr.Rectangle(1, 1, 10, 10)
		r, _ := cairo.RegionCreate()
		r.ContainsPoint(1, 5)
		da.QueueDrawRegion(r)
		da.AttachBuffers()
		fmt.Println(cr.GetCurrentPoint())
	})
	area.ShowAll()

	frag := characters.NewFrag(11, 11, 5, 30, 30, "./img/frag.png", game)
	frag.P.Render()
	usr := characters.NewHero(1, 1, 10, 5, 5, "./img/usr.png", game)
	usr.P.Render()

	win.Connect("key-press-event", func(tree *gtk.EventBox, ev *gdk.Event) {
		usr.TakeDamage(characters.Frag(*frag), func() {
			fmt.Println("TAKE DAMAGE!!!!", usr.Life)
		})
		usr.P.Render()
		btn := gdk.EventButtonNewFromEvent(ev)
		switch btn.State() {
		//up
		case 111:
		case 25:
			usr.P.StepUp()
		//down
		case 116:
		case 39:
			usr.P.StepDown()
		//left
		case 113:
		case 38:
			usr.P.StepLeft()
		//right
		case 114:
		case 40:
			usr.P.StepRight()
		default:
			usr.P.Render()
		}
	})
	win.Connect("button-press-event", func(tree *gtk.EventBox, ev *gdk.Event) bool {
		btn := gdk.EventButtonNewFromEvent(ev)
		switch btn.Button() {
		case gdk.BUTTON_PRIMARY:
			usr.Attack(int(btn.X()), int(btn.Y()), *weapon.NewArraw("./img/fireball.png"))
			return true
		case gdk.BUTTON_MIDDLE:
			fmt.Println("middle-click detected!")
			return true
		case gdk.BUTTON_SECONDARY:
			fmt.Println("right-click detected!")
			return true
		default:
			return false
		}
	})
	win.ShowAll()

	gtk.Main()
}
