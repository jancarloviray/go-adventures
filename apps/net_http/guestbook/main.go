package main

import (
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/fatih/structs"
	"html/template"
	"net/http"
	"os"
	"time"
)

const (
	DBdir = "db"
)

type Entries []*Entry

type Entry struct {
	ID        int
	Timestamp time.Time
	Name      string
	Message   string
}

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func main() {
	initialize()
	defer destroy()

	// routes
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)

	// server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, req *http.Request) {
	// local vars
	var query interface{}
	entries := Entries{}
	queryResults := make(map[int]struct{})

	// open database
	d, err := db.OpenDB(DBdir)

	if err != nil {
		panic(err)
	}

	// use collection
	docEntries := d.Use("Entries")

	// build query from json and convert to interface{}
	json.Unmarshal([]byte(`"all"`), &query)

	// execute query and pass results to queryResults
	if err := db.EvalQuery(query, docEntries, &queryResults); err != nil {
		panic(err)
	}

	// queryResults contains []int of IDs
	for id := range queryResults {
		entry := Entry{}

		readBack, _ := docEntries.Read(id)

		// map[string]interface{} TO struct hack

		j, _ := json.Marshal(readBack) // struct to json
		json.Unmarshal(j, &entry)      // json to actual type

		entries = append(entries, &entry)
	}

	// compile template with data
	if err := index.Execute(w, entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func add(w http.ResponseWriter, req *http.Request) {
	// make sure it's post
	if req.Method != "POST" {
		http.NotFound(w, req)
		return
	}

	entry := Entry{
		Name:    req.FormValue("name"),
		Message: req.FormValue("message"),
	}

	d, err := db.OpenDB(DBdir)
	if err != nil {
		panic(err)
	}

	docEntries := d.Use("Entries")

	docEntries.Insert(structs.New(entry).Map())

	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func initialize() {
	// remove temp files
	os.RemoveAll(DBdir)

	// open db
	d, err := db.OpenDB(DBdir)

	if err != nil {
		panic(err)
	}

	// create collection
	if err := d.Create("Entries"); err != nil {
		panic(err)
	}

	// collection instance
	docEntries := d.Use("Entries")

	// dummy data
	entries := Entries{
		&Entry{1, time.Now(), "Entry 1", "First Entry!"},
		&Entry{2, time.Now(), "Golang", "Go language is awesome"},
	}

	// insert each
	for _, entry := range entries {
		docEntries.Insert(structs.New(entry).Map())
	}
}

func destroy() {
	os.RemoveAll(DBdir)
}
