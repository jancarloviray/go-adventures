/*
Tying DB
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"os"
)

var dbDir string = "db"

func main() {
	setupDB()

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{Directory: "."}))
	m.Use(DB())

	m.Get("/tasks", func(r render.Render, db *db.DB) {
		r.HTML(200, "martini-02", GetAll(db))
	})

	m.Run()

	defer func() {
		os.RemoveAll(dbDir)
	}()
}

type dynamicType map[string]interface{}

func GetAll(db *db.DB) []dynamicType {
	arrayOfDynamicType := []dynamicType{}
	var dyna dynamicType

	tasks := db.Use("Tasks")

	_, err := tasks.Insert(map[string]interface{}{
		"Title": "Get Milk",
		"Done":  false,
	})

	tasks.Insert(map[string]interface{}{
		"Title": "Get Towel",
		"Done":  true,
	})

	if err != nil {
		panic(err)
	}

	tasks.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		json.Unmarshal([]byte(docContent), &dyna)
		arrayOfDynamicType = append(arrayOfDynamicType, dyna)
		fmt.Println(id, string(docContent), &dyna)
		return true
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(arrayOfDynamicType)

	return arrayOfDynamicType
}

// Middleware (will run on every request)
func DB() martini.Handler {
	// preparation
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

	if err := myDB.Close(); err != nil {
		panic(err)
	}
}
