package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

func main() {
	c := cron.New()

	entryID1, err := c.AddFunc("0,30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("01-entryId = %v\n", entryID1)

	entryID2, err := c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("02-entryId = %v\n", entryID2)

	entryID3, err := c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("03-entryId = %v\n", entryID3)

	c.Start()
}
