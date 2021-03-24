package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	// Do work here
	http.HandleFunc("/proxy.pac", func(w http.ResponseWriter, r *http.Request) {
		rAddr := r.RemoteAddr
		method := r.Method
		path := r.URL.Path
		log.Printf("%s [%s] %s\n", rAddr, method, path)
		http.ServeFile(w, r, "./proxy.pac")
	})

	log.Println("Listening...")
	http.ListenAndServe(":18080", nil)
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "PACServer",
		DisplayName: "Go PAC Server",
		Description: "local server for proxy auto-config",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err = s.Install()
			if err != nil {
				fmt.Print(err.Error())
				return
			}
		case "uninstall":
			err = s.Uninstall()
			if err != nil {
				fmt.Print(err.Error())
				return
			}
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Print(err.Error())
				return
			}
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Print(err.Error())
				return
			}

		}

	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

}
