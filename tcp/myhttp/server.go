package myhttp

import (
	"../../data"
	"io"
	"log"
	"net/http"
)

//type helloHandler struct{}

func getData(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, data.GET_DATA+"\n")
}

func sendData(w http.ResponseWriter, r *http.Request) {
	data.SEND_DATA = "wzn"
	io.WriteString(w, r.URL.Path)
}

func Start() {
	http.HandleFunc("/getData", getData)
	http.HandleFunc("/sendData", sendData)
	err := http.ListenAndServe(":10010", nil)

	log.Println("http listen port:10010")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
