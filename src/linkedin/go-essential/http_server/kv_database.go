/*
Key-value database

$ curl -d'hello' http://localhost:8080/k1
$ curl http://localhost:8080/k1
hello
$ curl -i http://localhost:8080/k2
404 not found

Value size is limited to 1k
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

var n int64 = 1 << 10

type Server struct {
	db DB
}

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// POST /key Store request body as value
// GET /<key> Send back value, or 404 if key not found
func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method, r.RequestURI)
	key := r.RequestURI[1:]
	switch r.Method {
	case "GET":
		dat := s.db.Get(key)
		if dat == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			fmt.Fprintln(w, string(dat))
		}
	case "POST":
		rd := bufio.NewReader(io.LimitReader(r.Body, n))
		defer r.Body.Close()
		dat, err := rd.ReadBytes(0)
		if err != io.EOF {
			log.Println(err)
			return
		}
		s.db.Set(key, dat)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	serv := Server{}
	http.HandleFunc("/health", serv.Health)
	http.HandleFunc("/", serv.Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
