package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/julienschmidt/httprouter"
	"github.com/v2af/ipify/api"
	"github.com/v2af/ipify/config"
)

func prepare() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	prepare()
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()
	handleConfig(*cfg)
}

func handleConfig(configFile string) {
	err := config.Parse(configFile)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", api.GetIP)

	router.NotFound = http.HandlerFunc(api.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(api.MethodNotAllowed)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config().Port), router))
}
