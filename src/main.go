package main

import (
	"fmt"
	"net/http"
)

func main() {
	CommandLineInfo.Init()
	InitConf(CommandLineInfo.ConfPath)

	fmt.Println("serve on port:", CommandLineInfo.Port)
	addr := fmt.Sprintf(":%d", CommandLineInfo.Port)
	http.ListenAndServe(addr, &HttpMoc{})
}
