package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var storage map[string]string = make(map[string]string) // TODO: Remove this

func handleSet(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var m map[string]string
	err = json.Unmarshal(body, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for k, v := range m {
		storage[k] = v
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	val, ok := storage[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, val)
}

type delReq struct {
	Key *string `json:"key"`
}

func handleDel(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var d delReq
	err = json.Unmarshal(body, &d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if d.Key == nil {
		http.Error(w, "JSON request must have field 'key'", http.StatusBadRequest)
		return
	}
	delete(storage, *d.Key)
}

func main() {
	http.HandleFunc("/set_key", handleSet)
	http.HandleFunc("/get_key", handleGet)
	http.HandleFunc("/del_key", handleDel)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	log.Fatal(http.ListenAndServe(":8089", nil))
}
