package main

import (
	"fmt"
	"os"
	"github.com/takama/daemon"
	"github.com/robfig/cron"

)

var (
	cronJob     *cron.Cron = cron.New()
	jobSchedule string     = "0 * * * * *"
)

func (service *Service) StartJob() (string, error) {
	f, _ := os.OpenFile("./logger.log", os.O_RDWR, 0777)
	f.WriteString("can write")

	cronJob.AddFunc(jobSchedule, func() {
	})

	cronJob.Start()
	return service.Start()
}

func main() {

	srv, err := daemon.New(name, description)
	service := &Service{srv}
	status, err := service.Manage()

	if err != nil {
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
