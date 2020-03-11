package main

import (
	"fmt"
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
			rep = rule.Response
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
