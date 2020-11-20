package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/byte/")
	resp, err := http.Get(fmt.Sprintf("http://httpbin.default.svc.cluster.local:8000/bytes/%s",id))
	if err != nil {
		fmt.Fprintf(w, "err is %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "err is %v", err)
		return
	}
	fmt.Println("get response ok")
	fmt.Fprintf(w, "%s", body)
	return
}

func main() {
	http.HandleFunc("/byte/", handler)
	fmt.Println("started...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}