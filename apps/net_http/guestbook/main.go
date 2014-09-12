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

/*
func Must(t *Template, err error) *Template
    Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as `var t = template.Must(template.New("name").Parse("html"))`

func ParseFiles(filenames ...string) (*Template, error)
    ParseFiles creates a new Template and parses the template definitions from the named files. The returned template's name will have the (base) name and (parsed) contents of the first file. There must be at least one file. If an error occurs, parsing stops and the returned *Template is nil.
*/

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

/*
func (t *Template) Execute(wr io.Writer, data interface{}) error
	Execute applies a parsed template to the specified data object, writing the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel.

func Error(w ResponseWriter, error string, code int)
    Error replies to the request with the specified error message and HTTP code. The error message should be plain text.
*/

func hello(w http.ResponseWriter, req *http.Request) {
	// QUERYING

	d, err := db.OpenDB(DBdir)

	if err != nil {
		panic(err)
	}

	docEntries := d.Use("Entries")

	// placeholder
	entries := Entries{}
	queryResults := make(map[int]struct{})

	// build query
	var query interface{}

	json.Unmarshal([]byte(`"all"`), &query)

	if err := db.EvalQuery(query, docEntries, &queryResults); err != nil {
		panic(err)
	}

	// res contains IDs
	for id := range queryResults {
		readBack, _ := docEntries.Read(id)
		entry := Entry{}
		j, _ := json.Marshal(readBack) // {"ID":2,"Message":"Go language is awesome","Name":"Golang","Timestamp":{}}
		json.Unmarshal(j, &entry)      // {2 0001-01-01 00:00:00 +0000 UTC Golang Go language is awesome}
		entries = append(entries, &entry)
	}

	if err := index.Execute(w, entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/*
func ListenAndServe(addr string, handler Handler) error
	ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Handler is typically nil, in which case the DefaultServeMux is used.
*/

func main() {
	initialize()
	defer destroy()

	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func initialize() {
	os.RemoveAll(DBdir)

	// OPEN DB (OR CREATE)

	d, err := db.OpenDB(DBdir)

	if err != nil {
		panic(err)
	}

	// CREATE COLLECTION

	if err := d.Create("Entries"); err != nil {
		panic(err)
	}

	// COLLECTION INSTANCE

	docEntries := d.Use("Entries")

	// INSERT

	entries := Entries{
		&Entry{1, time.Now(), "Entry 1", "First Entry!"},
		&Entry{2, time.Now(), "Golang", "Go language is awesome"},
	}

	for _, entry := range entries {
		docEntries.Insert(structs.New(entry).Map())
	}
}

func destroy() {
	os.RemoveAll(DBdir)
}
