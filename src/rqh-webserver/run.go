package main

import (
	"log"
	"flag"
	"github.com/mamemomonga/go-rqh-webserver/src/web"
)

func run() error {
	log.Printf("info: Basedir: %s", Basedir)
	log.Printf("info: rqh-webserver VERSION:%s REVISION: %s",Version,Revision)

	listen := flag.String("listen","127.0.0.1:8000","Listen host:port")
	flag.Parse()

	w := web.New()
	err := w.Run(*listen)
	if err != nil {
		return err
	}
	return nil
}

