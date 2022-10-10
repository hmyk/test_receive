package main

import (
	"fmt"
	"io"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
		<!DOCTYPE html>
		<html lang="ja">
		<head>
			<meta charset="UTF-8">
			<title>確認サイト</title>
		</head>
		<body>
			<h1>state受信OK</h1>
			<p>コード受信しました</p>
		</body>
		</html>`)
}

func main() {
	fmt.Printf("main start")

	http.HandleFunc("/", handler)

	http.ListenAndServe("localhost:****", nil)

}
