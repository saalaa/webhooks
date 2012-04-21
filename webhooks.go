// Copyright 2012 (C) Benoit Myard <benoit@saalaa.net>
// All rights reserved.

package main

import (
	"flag"
	"net/http"
	"os/exec"
	"log"
)

func main () {
	httpInterface := *flag.String("i", "127.0.0.1:8005", "Interface to listen to.")

	http.HandleFunc("/", func (res http.ResponseWriter, req *http.Request) {
		var code int
		var command *exec.Cmd

		commandPath := "." + req.URL.Path

		path, err := exec.LookPath(commandPath)
		if err != nil {
			code = http.StatusNotImplemented
			goto exit
		}

		command = exec.Command(path)

		if len(req.URL.Query().Get("async")) > 0 {
			err := command.Start()

			if err != nil {
				code = http.StatusInternalServerError
			}
		} else {
			err := command.Run()

			if err != nil {
				code = http.StatusInternalServerError
			}
		}

	exit:
		log.Println(req.URL.Path, path, code)
		res.WriteHeader(code)
	})

	flag.Parse()

	err := http.ListenAndServe(httpInterface, nil)
	if err != nil {
		panic("The HTTP server could not be started")
	}
}
