package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"strconv"
	//"html"
)

type tomlConfig struct {
	Port  int
	Relay map[string]relayInfo
}

type relayInfo struct {
	Path     string
	EndPoint string
}

var config tomlConfig

func topHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("%s", r.URL.Path)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("Err: %s", err)
	}

	for h := range r.Header {
		for v := range r.Header[h] {
			glog.Infof("Header: %s: %s", h, r.Header[h][v])
		}
	}
	glog.Infof("Method: %s", r.Method)
	glog.Infof("Path: %s", r.URL.Path)
	glog.Infof("Body: %s", body)

	fmt.Fprintf(w, "OK\n")
	var endpoint string
	for e := range config.Relay {
		if r.URL.Path == config.Relay[e].Path {
			endpoint = config.Relay[e].EndPoint
			break
		}
	}
	glog.Infof("EndPoint: %s", endpoint)
	if endpoint == "" {
		return
	}

	req, err := http.NewRequest(r.Method, endpoint, bytes.NewReader(body))
	if err != nil {
		glog.Errorf("Err: %s", err)
	}
	req.Header = r.Header
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Errorf("Err: %s", err)
	}
	defer resp.Body.Close()
	relay_body, err := ioutil.ReadAll(resp.Body)
	glog.Infof("Relay Body: %s", relay_body)
}

func main() {
	var conffile = flag.String("conf", "", "Configuration File")
	flag.Parse()
	glog.Infof("Config file: %s", *conffile)

	if *conffile != "" {
		if _, err := toml.DecodeFile(*conffile, &config); err != nil {
			glog.Fatalf("Error on decoding config file: %s", err)
		}
	}

	port := config.Port
	if port == 0 {
		port = 18080
	}
	http.HandleFunc("/", topHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
	glog.Error("Shutting down..")
}
