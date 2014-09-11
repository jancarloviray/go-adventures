/*
Templating and Rendering
*/

package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	// render middleware
	// https://github.com/codegangsta/martini-contrib/tree/master/render
	m.Use(render.Renderer(render.Options{
		Directory: ".",
	}))

	// This will set the Content-Type header to "text/html; charset=UTF-8"
	m.Get("/wishes", func(r render.Render) {
		r.HTML(200, "martini-02", nil)
	})

	// This will set the Content-Type header to "text/html; charset=UTF-8"
	m.Get("/text", func(r render.Render) {
		r.HTML(200, "hello", "world")
	})

	// This will set the Content-Type header to "application/json; charset=UTF-8"
	m.Get("/json", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	m.Run()
}
