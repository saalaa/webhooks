// Copyright 2012 (C) Benoit Myard <benoit@saalaa.net>
// All rights reserved.

package main

import (
	"flag"
	"net/http"
	"os/exec"
	"log"
	"io"
)

func main () {
	httpInterface := *flag.String("i", "127.0.0.1:8005", "Interface to listen to."

	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		func reply(code int) {
			if code == http.StatusOK {
				log.Error(req.URL.Path, code)
			} else {
				log.Info(req.URL.Path, code)
			}

			res.WriteHeader(code)
		}

		path, err := exec.LookPath(command)
		if err != nil {
			return reply(http.StatusNotImplemented)
		}

		command := exec.Command(path)

		if len(req.URL.Query().Get("async")) > 0 {
			err := command.Start()
		} else {
			err := command.Run()
		}

		if err != nil {
			return reply(http.StatusInternalServerError)
		}

		reply(http.StatusOK)
	})

	err := http.ListenAndServe(httpInterface, nil)
	if err != nil {
		panic("The HTTP server could not be started")
	}
}
