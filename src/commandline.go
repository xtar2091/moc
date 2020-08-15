package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

type CommandLine struct {
	ConfPath string
	Port     int
}

var CommandLineInfo = &CommandLine{}

func (obj *CommandLine) Init() {
	fs := FileSystem{}
	homeDir := getHomeDir()
	homeConfPath := homeDir + string(os.PathSeparator) + ".moc" + string(os.PathSeparator) + "conf.json"
	CommandLineInfo.ConfPath = ""
	flag.Parse()
	if flag.NArg() == 1 {
		CommandLineInfo.ConfPath = flag.Arg(0)
	} else if fs.IsFile("conf.json") {
		CommandLineInfo.ConfPath = "conf.json"
	} else if fs.IsFile(homeConfPath) {
		CommandLineInfo.ConfPath = homeConfPath
	} else if fs.IsFile("/etc/moc/conf.json") {
		CommandLineInfo.ConfPath = "/etc/moc/conf.json"
	} else {
		log.Fatalln("conf.json not found")
	}
	log.Println("user conf:", CommandLineInfo.ConfPath)
}

func getHomeDir() string {
	user, err := user.Current()
	if err == nil {
		return user.HomeDir
	}

	if runtime.GOOS == "widows" {
		return getWindowsHomeDir()
	}

	return getUnixHomeDir()
}

func getWindowsHomeDir() string {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}

	return home
}

func getUnixHomeDir() string {
	// First prefer the HOME environmental variable
	home := os.Getenv("HOME")
	if home != "" {
		return home
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		log.Println("get unix home dir failed, error:", err)
		return ""
	}

	result := strings.TrimSpace(stdout.String())
	return result
}
