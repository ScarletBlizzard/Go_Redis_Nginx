package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx context.Context = context.Background()
var rc *redis.Client

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
		err = rc.Set(ctx, k, v, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	val, err := rc.Get(ctx, key).Result()
	if err != nil {
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
	err = rc.Del(ctx, *d.Key).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	rc = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB
	})

	http.HandleFunc("/set_key", handleSet)
	http.HandleFunc("/get_key", handleGet)
	http.HandleFunc("/del_key", handleDel)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	log.Fatal(http.ListenAndServe(":8089", nil))
}
