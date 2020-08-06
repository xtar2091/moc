package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type HttpMoc struct {

}

func (obj *HttpMoc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var responseBytes []byte
	defer func() {
		w.Write(responseBytes)
	}()

	log.Println("new request received, method:", r.Method, ", path:", r.URL.Path)
	key := MakeMocKey(r.Method, r.URL.Path)
	moc, ok := GlobalConf[key]
	if !ok {
		responseBytes = []byte("welcome to moc")
		log.Println("unknown method or path")
		return
	}

	if moc.Sleep > 0 {
		log.Println("sleep ", moc.Sleep, "milliseconds")
		time.Sleep(time.Duration(moc.Sleep) * time.Millisecond)
	}

	queryString := r.URL.RawQuery
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseBytes = []byte("welcome to moc")
		log.Println("read request body failed, error:", err)
		return
	}
	filter := RulesFilter{}
	responseBytes = []byte(filter.DoFilter(queryString, string(body), moc.Rules))
}
