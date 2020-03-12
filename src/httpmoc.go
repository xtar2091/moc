package main

import (
	"fmt"
	"io/ioutil"
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

	key := MakeMocKey(r.Method, r.URL.Path)
	moc, ok := GlobalConf[key]
	if !ok {
		responseBytes = []byte("welcome to moc")
		return
	}

	if moc.Sleep > 0 {
		time.Sleep(time.Duration(moc.Sleep) * time.Millisecond)
	}

	queryString := r.URL.RawQuery
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseBytes = []byte("welcome to moc")
		fmt.Println(err)
		return
	}
	filter := RulesFilter{}
	responseBytes = []byte(filter.DoFilter(queryString, string(body), moc.Rules))
}
