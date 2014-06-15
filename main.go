package main

import (
  "github.com/BurntSushi/toml"
  "flag"
  "log"
  "net/http"
  "fmt"
  "io/ioutil"
  //"html"
)

type tomlConfig struct {
  Auth authInfo
  Webhook webhookInfo
}

type authInfo struct {
  Id string
  Password string
}

type webhookInfo struct {
  Name string
  Type string
  EndPoint string
  Format string
}

func topHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("%s", r.URL.Path)
  body, err := ioutil.ReadAll(r.Body);
  if err != nil {
    log.Printf("%s", err)
  }

  log.Printf("%s", body)

  fmt.Fprintf(w, "OK")
}

func main() {
  var conffile = flag.String("conf", "", "Configuration File")
  flag.Parse()

  log.Printf("Config file: %s", *conffile)

  var config tomlConfig
  if *conffile != "" {
    if _, err := toml.DecodeFile(*conffile, &config); err != nil {
      panic(err)
    }
  }

  log.Print(config.Webhook.Name)

  http.HandleFunc("/", topHandler)
  http.ListenAndServe(":18080", nil)
  log.Print("Shutting down..")
}
