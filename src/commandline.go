package main

import "flag"

type CommandLine struct {
	ConfPath string
	Port int
}

var CommandLineInfo = &CommandLine{}

func (obj *CommandLine) Init() {
	flag.StringVar(&obj.ConfPath, "conf", "./conf.json", "config file path")
	flag.IntVar(&obj.Port, "port", 12345, "port of the server")
	flag.Parse()
}
