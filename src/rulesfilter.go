package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

type RulesFilter struct {

}

func (obj RulesFilter) DoFilter(queryString, body string, rules []Rules) string {
	rep := "welcome to moc"
	for _, rule := range rules {
		queryStringMatch := obj.doMatch(queryString, rule.Request)
		bodyMatch := obj.doMatch(body, rule.Body)
		if queryStringMatch && bodyMatch {
			if len(rule.Response) > 0 {
				rep = rule.Response
				log.Println("response from conf")
			} else if len(rule.ResponseFile) > 0 {
				rep = obj.readResponseFromFile(rule.ResponseFile)
				log.Println("response from file:", rule.ResponseFile)
			} else if len(rule.ResponseShell) > 0 {
				rep = obj.readResponseFromShell(rule.ResponseShell, queryString, body)
				log.Println("response from shell:", rule.ResponseShell)
			}
			break
		}
	}
	return rep
}

func (obj RulesFilter) doMatch(text, regexString string) bool {
	if (text == "") || (regexString == "") {
		return true
	}
	reg, err := regexp.Compile(regexString)
	if err != nil {
		log.Println("regex match failed, error:", err)
		return false
	}
	return reg.MatchString(text)
}

func (obj RulesFilter) readResponseFromFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("open file failed, file name:", fileName)
		return "welcome to moc"
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	rep, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Println("read file failed, file name:", fileName, ", error:", err)
		return "welcome to moc"
	}
	return string(rep)
}

func (obj RulesFilter) readResponseFromShell(shellFileName, queryString, body string) string {
	ext := filepath.Ext(shellFileName)
	if ext == ".py" {
		cmd := exec.Command("python", shellFileName, queryString, body)
		buf, err := cmd.Output()
		if err != nil {
			log.Println("read shell output failed, error:", err)
			return "welcome to moc"
		}
		return string(buf)
	}
	log.Println("unsupported shell:", shellFileName)
	return "welcome to moc"
}
