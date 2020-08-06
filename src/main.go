package main

import (
	"fmt"
	"log"
	"net/http"
)

func initLog() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	initLog()
	CommandLineInfo.Init()
	InitConf(CommandLineInfo.ConfPath)

	log.Println("serve on port:", CommandLineInfo.Port)
	addr := fmt.Sprintf(":%d", CommandLineInfo.Port)
	http.ListenAndServe(addr, &HttpMoc{})
}
