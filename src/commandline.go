package main

import (
	"flag"
	"log"
)

type CommandLine struct {
	ConfPath string
	Port int
}

var CommandLineInfo = &CommandLine{}

func (obj *CommandLine) Init() {
	fs := FileSystem{}
	flag.Parse()
	if flag.NArg() == 1 {
		CommandLineInfo.ConfPath = flag.Arg(0)
	} else if fs.IsFile("conf.json") {
		CommandLineInfo.ConfPath = "conf.json"
	} else if fs.IsFile("~/.moc/conf.json") {
		CommandLineInfo.ConfPath = "~/.moc/conf.json"
	} else if fs.IsFile("/etc/moc/conf.json") {
		CommandLineInfo.ConfPath = "/etc/moc/conf.json"
	} else {
		log.Fatalln("conf.json not found")
	}
}
