package jobs

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

var scheduler *gocron.Scheduler

func init() {
	fixedZone, _ := time.LoadLocation("Asia/Jakarta")
	scheduler = gocron.NewScheduler(fixedZone)
	fmt.Printf(">>> Initiate jobs success \n")

	scheduleList()
	scheduler.StartAsync()
}

// scheduleList ..
func scheduleList() {
	// Segame
	//job, err := scheduler.Every(1).Day().At("16:00").Do(Segame)
	//log.Printf("\nJob: %v, Error: %v \n", job, err)
}
