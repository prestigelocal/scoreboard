package main

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"github.com/treetopllc/elastilog"
)

func mlbPing() {

	c := elastilog.NewClient("http://172.17.0.3:9200", "ping", "scoreboard")
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

	c.Send(elastilog.Entry{
		Timestamp: time.Now(),
		Host:      "debian-jessie",
		Log:       "SLOW:",
		Attributes: elastilog.Attributes{
			"payload":         string(body),
		},
	})
	c.Close()
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	if len(body) != 0 {
	}
}

