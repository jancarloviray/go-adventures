/*
Embedded Database
to install:
go get "github.com/HouzuoGuo/tiedot/db"
go get "github.com/codegangsta/martini"
go get "github.com/codegangsta/martini-contrib/binding"
go get "github.com/codegangsta/martini-contrib/render"
go run martini-03.go
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"os"
)

var dbDir string = "db"

func main() {
	m := martini.Classic()

	setup()
	middlewares(m)
	routes(m)

	m.Run()
	defer cleanup()
}

func setup() {
	setupDB()
}

func middlewares(m *martini.ClassicMartini) {
	m.Use(render.Renderer(render.Options{Directory: "."}))
	m.Use(DB())
}

func routes(m *martini.ClassicMartini) {
	m.Get("/tasks", func(r render.Render, db *db.DB) {
		r.HTML(200, "martini-02", GetAll(db))
	})
	m.Post("/tasks", binding.Form(Task{}), func(task Task, r render.Render, db *db.DB) {
		CreateTask(db, &task)
		r.HTML(200, "martini-02", GetAll(db))
	})
}

func cleanup() {
	os.RemoveAll(dbDir)
}

// DATA STRUCTURES

type Tasks []Task

type Task struct {
	Title string `form:"title" binding:"required"`
	Done  bool   `form:"done"`
}

// ROUTE HANDLERS

func GetAll(db *db.DB) Tasks {
	tasks, t := Tasks{}, Task{}
	db.Use("Tasks").ForEachDoc(func(id int, body []byte) (next bool) {
		json.Unmarshal(body, &t)
		tasks = append(tasks, t)
		return true
	})
	return tasks
}

func CreateTask(db *db.DB, task *Task) {
	// why this 3-line mess? because interface of collection.Insert is this:
	// func (col *Col) Insert(doc map[string]interface{}) (id int, err error) { }
	jsonTask := map[string]interface{}{}
	j, _ := json.Marshal(task)
	json.Unmarshal(j, &jsonTask)

	dbTasks := db.Use("Tasks")
	dbTasks.Insert(jsonTask)
}

// MIDDLEWARES

func DB() martini.Handler {
	myDB, err := db.OpenDB(dbDir)
	if err != nil {
		panic(err)
	}
	// return will be called on every request
	return func(c martini.Context) {
		// defer myDB.Close() // for some reason, this prevents connection from completing

		// maps an instance of database to our request context
		// this allows subsequent handler functions to specify
		// this particular db type as an argument and get it
		// (automagically) injected
		c.Map(myDB)
		c.Next()
	}

}

// SETUP

func setupDB() {
	os.RemoveAll(dbDir)
	myDB, err := db.OpenDB(dbDir)
	if err != nil {
		panic(err)
	}
	if err := myDB.Create("Tasks"); err != nil {
		panic(err)
	}
	if err := myDB.Create("Users"); err != nil {
		panic(err)
	}
	for key, name := range myDB.AllCols() {
		fmt.Println("Existing Collection:", key, name)
	}
	setupDummy(myDB)
	if err := myDB.Close(); err != nil {
		panic(err)
	}
}

func setupDummy(db *db.DB) {
	dbTasks := db.Use("Tasks")
	dummyTasks := []map[string]interface{}{
		map[string]interface{}{"title": "Get Milk", "done": true},
		map[string]interface{}{"title": "Drink Milk", "done": false},
	}
	for _, t := range dummyTasks {
		if _, err := dbTasks.Insert(t); err != nil {
			panic(err)
		}
	}
}
