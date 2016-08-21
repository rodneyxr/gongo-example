package main

import (
	"github.com/flosch/pongo2"
	"github.com/rodneyxr/gongo"
)

func index(r *gongo.Request) {
	context := pongo2.Context{
		"title":   "Home",
		"message": "Hello, World!",
	}
	gongo.Render(r, "index.html", context)
}

func page1(r *gongo.Request) {
	gongo.Render(r, "page1.html", nil)
}

func hello(r *gongo.Request) {
	name := r.Params.ByName("name")
	context := pongo2.Context{
		"name": name,
	}
	gongo.Render(r, "hello.html", context)
}
