package main

import "gtk/cmd/person"

func IndexObject(person ...person.IPerson) {
	for _, object := range person {
		object.Render()
	}
}
