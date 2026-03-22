package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/ikawaha/kagome-dict/uni"
	"github.com/mattn/go-haiku"
)

type Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

// JSONエンコーダ再利用（GC削減）
var bufPool = sync.Pool{
	New: func() any {
		return make([]byte, 0, 512)
	},
}

func writeJSON(w http.ResponseWriter, status int, v Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	buf := bufPool.Get().([]byte)
	buf = buf[:0]

	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, `{"error":"encode error"}`, 500)
		return
	}

	w.Write(data)

	bufPool.Put(buf)
}

// POST /match
// body: { "text": "...", "rule": [5,7,5] }
func matchHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
		Rule []int  `json:"rule"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, Response{Error: "invalid json"})
		return
	}

	ok := haiku.Match(req.Text, req.Rule)

	writeJSON(w, 200, Response{
		Result: ok,
	})
}

// POST /find
// body: { "text": "...", "rule": [5,7,5] }
func findHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
		Rule []int  `json:"rule"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, 400, Response{Error: "invalid json"})
		return
	}

	result := haiku.Find(req.Text, req.Rule)

	writeJSON(w, 200, Response{
		Result: result,
	})
}

func main() {
	haiku.UseDict(uni.Dict())

	mux := http.NewServeMux()

	mux.HandleFunc("/match", matchHandler)
	mux.HandleFunc("/find", findHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("server started :8080")
	log.Fatal(server.ListenAndServe())
}
