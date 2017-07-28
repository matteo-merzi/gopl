package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/matteo-merzi/gopl/ch07/7.8"
)

var tracks = []*column.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, column.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, column.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, column.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, column.Length("4m24s")},
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := column.NewByColumns(tracks)
	switch r.FormValue("sort") {
	case "title":
		c.Select(column.ByTitle)
	case "artist":
		c.Select(column.ByArtist)
	case "album":
		c.Select(column.ByAlbum)
	case "year":
		c.Select(column.ByYear)
	case "length":
		c.Select(column.ByLength)
	}
	sort.Sort(c)
	// column.PrintTracks(tracks)
	err := tpl.Execute(w, tracks)
	if err != nil {
		log.Printf("template error: %s", err)
	}
}
