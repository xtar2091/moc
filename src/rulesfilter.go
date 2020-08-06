package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
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
			} else if len(rule.ResponseFile) > 0 {
				rep = obj.readResponseFromFile(rule.ResponseFile)
			} else if len(rule.ResponseShell) > 0 {
				rep = obj.readResponseFromShell(rule.ResponseShell)
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
		fmt.Println(err)
		return false
	}
	return reg.MatchString(text)
}

func (obj RulesFilter) readResponseFromFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return "welcome to moc"
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	rep, err := ioutil.ReadAll(reader)
	if err != nil {
		return "welcome to moc"
	}
	return string(rep)
}

func (obj RulesFilter) readResponseFromShell(fileName string) string {
	return ""
}
