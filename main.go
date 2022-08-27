package main

import (

	//"time"
	"cron-go/models"

	"github.com/jasonlvhit/gocron"
)

func main() {

	// main cron
	gocron.Every(5).Second().Do(models.SchedulerHarian)
	//gocron.Every(1).Day().At("16:27").Do(models.SchedulerHarian)
	<-gocron.Start()

}
