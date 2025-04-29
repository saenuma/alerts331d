package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	serverKey = "p3WuLMz8T1m34b9yaYC3kq9yRibC0YrTT60Nh4qn8mX4xbzmw2ig5J1AGwSJgrVxaQ2xLggemTgDvURrrPHcSJNqzcY5520MBKTe"
	addr      = "http://sae.ng"
)

type Msg struct {
	Subject, Body, ViewAddr string
	UnixTimestamp           int
}

func main() {
	msgs := []Msg{
		{"Home today", "Home Today", addr, 1745625600},
		{"Home today Again", "Home Today Yeah", addr, 1745625600},
		{"Office today", "Yeah office today", addr, 1745625200},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("k") == serverKey {
			// allow
			jsonBytes, _ := json.Marshal(msgs)
			w.Write(jsonBytes)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "not ok")
		}

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8086"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
