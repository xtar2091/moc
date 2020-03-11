package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type Conf struct {
	Moc []Moc `json:"moc"`
}

type Rules struct {
	Body string `json:"body"`
	Request string `json:"request"`
	Response string `json:"response"`
}

type Moc struct {
	Path string `json:"path"`
	Method string `json:"method"`
	Sleep int64 `json:"sleep"`
	Rules []Rules `json:"rules"`
}

var GlobalConf map[string]Moc

func InitConf(confPath string) {
	content, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	conf := Conf{}
	err = json.Unmarshal(content, &conf)
	if err != nil {
		panic(err)
	}

	GlobalConf = make(map[string]Moc)
	for _, moc := range conf.Moc {
		key := MakeMocKey(moc.Method, moc.Path)
		GlobalConf[key] = moc
	}
}

func MakeMocKey(method, path string) string {
	return strings.ToLower(method) + strings.ToLower(path)
}
