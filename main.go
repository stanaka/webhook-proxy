package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	//"html"
)

type tomlConfig struct {
	Port  int
	Relay relayInfo
}

type relayInfo struct {
	Name     string
	Type     string
	EndPoint string
}

var config tomlConfig

func topHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s", r.URL.Path)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err: %s", err)
	}

	for h := range r.Header {
		for v := range r.Header[h] {
			log.Printf("Header: %s: %s", h, r.Header[h][v])
		}
	}
	log.Printf("Body: %s", body)
	log.Printf("EndPoint: %s", config.Relay.EndPoint)
	log.Printf("Method: %s", r.Method)

	fmt.Fprintf(w, "OK\n")
	if config.Relay.EndPoint == "" {
		return
	}

	req, err := http.NewRequest(r.Method,
		config.Relay.EndPoint, bytes.NewReader(body))
	if err != nil {
		log.Printf("Err: %s", err)
	}
	req.Header = r.Header
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Err: %s", err)
	}
	defer resp.Body.Close()
	relay_body, err := ioutil.ReadAll(resp.Body)
	log.Printf("Relay Body: %s", relay_body)
}

func main() {
	var conffile = flag.String("conf", "", "Configuration File")
	flag.Parse()

	log.Printf("Config file: %s", *conffile)

	if *conffile != "" {
		if _, err := toml.DecodeFile(*conffile, &config); err != nil {
			panic(err)
		}
	}

	port := config.Port
	if port == 0 {
		port = 18080
	}
	http.HandleFunc("/", topHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
	log.Print("Shutting down..")
}
