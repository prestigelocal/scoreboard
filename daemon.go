package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/takama/daemon"
	"fmt"
	"net/http"
)

const (
	name        = "myservice"
	description = "My Echo Service"
)

type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	http.HandleFunc("/mlb", func(w http.ResponseWriter, r *http.Request) {

		mlbPing(w, r)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test")
	})

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		fmt.Println("Server error")
		os.Exit(0)
	}

	return usage, nil
}