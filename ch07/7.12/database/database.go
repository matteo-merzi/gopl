package database

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/list.gohtml"))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type MutexDB struct {
	sync.Mutex
	db map[string]dollars
}

func (m *MutexDB) Create(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	if _, ok := m.db[item]; ok {
		msg := fmt.Sprintf("%q already exists\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	priceStr := req.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		http.Error(w, "no valid price given", http.StatusBadRequest)
		return
	}

	m.Lock()
	if m.db == nil {
		m.db = make(map[string]dollars, 0)
	}
	m.db[item] = dollars(price)
	m.Unlock()
}

func (m *MutexDB) Read(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	if _, ok := m.db[item]; !ok {
		msg := fmt.Sprintf("%q doesn't exists\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	m.Lock()
	fmt.Fprintf(w, "%s: %s\n", item, m.db[item])
	m.Unlock()
}

func (m *MutexDB) Update(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	if _, ok := m.db[item]; !ok {
		msg := fmt.Sprintf("%q doesn't exists\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	priceStr := req.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		http.Error(w, "no valid price given", http.StatusBadRequest)
		return
	}

	m.Lock()
	m.db[item] = dollars(price)
	m.Unlock()
}

func (m *MutexDB) Delete(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	if item == "" {
		http.Error(w, "no item given", http.StatusBadRequest)
		return
	}

	if _, ok := m.db[item]; !ok {
		msg := fmt.Sprintf("%q doesn't exists\n", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	m.Lock()
	delete(m.db, item)
	m.Unlock()
}

func (m *MutexDB) List(w http.ResponseWriter, req *http.Request) {
	m.Lock()
	tpl.Execute(w, m.db)
	m.Unlock()
}

func (m *MutexDB) Price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	m.Lock()
	if price, ok := m.db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	m.Unlock()
}
