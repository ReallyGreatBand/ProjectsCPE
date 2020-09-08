package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Session struct {
	Time string `json:"time"`
}

func MustEncode(enc *json.Encoder, v Session) {
	err := enc.Encode(v)
	if err != nil {
		panic(err)
	}
}

func jsonHandler(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	t := time.Now().Format(time.RFC3339)
	showTime := Session{
		Time: t,
	}
	MustEncode(json.NewEncoder(w), showTime)
}

func main() {
	http.HandleFunc("/time", jsonHandler)
	log.Fatal(http.ListenAndServe(":8795", nil))
}
