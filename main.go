package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"blue": "#0000FF", "red": "#ff0000"}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(db.list))
	mux.Handle("/hex", http.HandlerFunc(db.hex))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]string

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for color, hex := range db {
		fmt.Fprintf(w, "%s: %s\n", color, hex)
	}
}

func (db database) hex(w http.ResponseWriter, req *http.Request) {
	color := req.URL.Query().Get("color")
	hex, ok := db[color]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such color: %q\n", color)
		return
	}
	fmt.Fprintf(w, "%s\n", hex)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	color := req.URL.Query().Get("color")
	hex := req.URL.Query().Get("hex")
	db[color] = hex
	fmt.Fprintf(w, "created %s: %s\n", color, hex)

}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	color := req.URL.Query().Get("color")
	delete(db, color)
	fmt.Fprintf(w, "deleted %s\n", color)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	color := req.URL.Query().Get("color")
	hex := req.URL.Query().Get("hex")
	_, ok := db[color]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such color: %q\n", color)
		return
	}
	db[color] = hex
	fmt.Fprintf(w, "updated %s: %s\n", color, hex)
}
