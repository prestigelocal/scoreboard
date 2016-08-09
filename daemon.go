package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"github.com/takama/daemon"
	"fmt"
	"net/http"
	"io/ioutil"
)

const (
	name        = "myservice"
	description = "My Echo Service"
	port = ":8081"
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

	listener, err := net.Listen("tcp", port)
	if err != nil {
		return "Possibly was a problem with the port binding", err
	}

	listen := make(chan net.Conn, 100)
	go acceptConnection(listener, listen)

	for {
		select {
		case conn := <-listen:
			go handleClient(conn)
		case killSignal := <-interrupt:
			fmt.Println("Got signal:", killSignal)
			fmt.Println("Stoping listening on ", listener.Addr())
			listener.Close()
			if killSignal == os.Interrupt {
				return "Daemon was interruped by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}

	http.HandleFunc("/mlb", func(w http.ResponseWriter, r *http.Request) {

		esClient := initESClient()
		indexName := "mlb"
		indexExists, err := esClient.IndexExists(indexName).Do()
		if !indexExists {
			resp, err := esClient.CreateIndex(indexName).Do()
			if !resp.Acknowledged {
				fmt.Println(resp, "\nError: ", err)
			}
			fmt.Printf("Index %s created", indexName)

		} else {
			fmt.Printf("Index %s already created", indexName)
		}

		url := "http://gd2.mlb.com/components/game/mlb/year_2016/month_08/day_08/master_scoreboard.json"
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error", err.Error())
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println("response", body)
		if err != nil {
			fmt.Println("error", err.Error())
			return
		}
		if len(body) != 0 {
			resp, err := IndexDocJSONBytes(esClient, indexName, "document", string(body))
			if err != nil {
				fmt.Println(resp, "\nError: ", err)
			}
			fmt.Println(resp)
		}
	})

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error")
		os.Exit(0)
	}

	return usage, nil
}

func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		listen <- conn
	}
}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		client.Write(buf[:numbytes])
	}
}