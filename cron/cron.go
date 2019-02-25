package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main(){
	fmt.Println("Cron Starting...")
	//logging.Info("Cron Starting...")

	c := cron.New()

	c.AddFunc("* * * * * *",func(){
		log.Println("Cron Run AddFunc-1")
	})

	c.AddFunc("1 * * * * *",func(){
		log.Println("Cron Run AddFunc-2")
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select{
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
