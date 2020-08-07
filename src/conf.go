package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

type Conf struct {
	// moc规则集
	Moc []Moc `json:"moc"`
	Port int `json:"port"`
	// 非必填。指定root_path后，response_file、response_shell 可使用相对路径
	ResponseRootPath string `json:"response_root_path"`
}

type Rules struct {
	// 非必填。post请求body正则表达式，若不填或为空视为匹配成功
	Body string `json:"body"`
	// 非必填。查询字符串正则表达式，url"?"后面的内容，若不填或为空视为匹配成功
	Request string `json:"request"`
	// 非必填。返回的响应数据
	Response string `json:"response"`
	// 非必填。响应返回指定文件中的内容
	ResponseFile string `json:"response_file"`
	// 非必填。获取指定脚本的标准输出并返回
	ResponseShell string `json:"response_shell"`
}

type Moc struct {
	// url的path部分
	Path string `json:"path"`
	// 请求的方法
	Method string `json:"method"`
	// 非必填。休眠一段时间再向客户端返回，单位：毫秒
	Sleep int64 `json:"sleep"`
	// 规则集
	Rules []Rules `json:"rules"`
}

var GlobalConf map[string]Moc

func InitConf(confPath string) {
	content, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalln("read conf file failed, conf file:", confPath, ", error:", err)
	}

	conf := Conf{}
	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatalln("parse conf file to json failed, conf file:", confPath, ", error:", err)
	}

	CommandLineInfo.Port = conf.Port

	fs := FileSystem{}
	GlobalConf = make(map[string]Moc)
	for _, moc := range conf.Moc {
		for i := 0; i < len(moc.Rules); i++ {
			file := filepath.Join(conf.ResponseRootPath, moc.Rules[i].ResponseFile)
			file = filepath.Clean(file)
			if fs.IsFile(file) == true {
				moc.Rules[i].ResponseFile = file
			} else {
				moc.Rules[i].ResponseFile = ""
			}

			file = filepath.Join(conf.ResponseRootPath, moc.Rules[i].ResponseShell)
			file = filepath.Clean(file)
			if fs.IsFile(file) == true {
				moc.Rules[i].ResponseShell = file
			} else {
				moc.Rules[i].ResponseShell = ""
			}
		}
		key := MakeMocKey(moc.Method, moc.Path)
		GlobalConf[key] = moc
	}
}

func MakeMocKey(method, path string) string {
	key := strings.ToLower(method) + strings.ToLower(path)
	return key
}
