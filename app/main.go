package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/sleep", sleep)
	http.HandleFunc("/post", post)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	res := []byte(`<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
</head>
<body style="margin: 100px;">
	<a href="https://github.com/pokutuna/collect-browser-request-cancellation-error">pokutuna&#x2F;collect-browser-request-cancellation-error</a
	<ul>
		<li><a href="/">/<a></li>
		<li><a href="/sleep">/sleep<a></li>
	</ul>
</body>
</html>
`)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html")
	w.Write(res)
}

func sleep(w http.ResponseWriter, r *http.Request) {
	sec, err := strconv.ParseInt(r.URL.Query().Get("s"), 10, 64)
	if err != nil {
		sec = 5
	}

	time.Sleep(time.Second * time.Duration(sec))

	res, _ := json.Marshal(map[string]int64{"sleep": sec})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
	}

	io.Copy(os.Stdout, r.Body)
	fmt.Print("\n")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("ok"))
}
