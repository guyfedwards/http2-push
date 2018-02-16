package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var image []byte

func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

// send html and push image
func handleHTML(w http.ResponseWriter, r *http.Request) {
	pusher, ok := w.(http.Pusher)
	if ok {
		fmt.Println("Push image")
		pusher.Push("/image", nil)
	}
	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	cert := flag.String("cert", "", "cert for ssl")
	key := flag.String("key", "", "key for ssl")
	port := flag.Int("port", 443, "port for ssl")

	flag.Parse()

	http.HandleFunc("/", handleHTML)
	http.HandleFunc("/image", handleImage)
	fmt.Println("start http listening :18443")

	err := http.ListenAndServeTLS(":"+strconv.Itoa(*port), *cert, *key, nil)

	fmt.Println(err)
}
