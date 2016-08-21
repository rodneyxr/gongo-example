package main

import "github.com/rodneyxr/gongo"

// Routes sets up the routes for gongo
func Routes(r *gongo.Router) {
	r.GET("/", index)
	r.GET("/page1", page1)
	r.GET("/hello/:name", hello)
}
