package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/benjimouse/timelogutil"
	"github.com/tkanos/gonfig"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	currentTime := time.Now() // The current time
	configuration := timelogutil.Configuration{}

	// TODO: something with the file name to cope with dev / prod! environments
	gonfig.GetConf("config/config.development.json", &configuration)

	session := timelogutil.GetMongoSession(configuration)
	defer session.Close()

	c := session.DB(configuration.Database).C("events")

	// Really if you're not going to send in an argument then ...
	if len(os.Args) != 1 {
		a := os.Args[1] // The value passed into the command line

		task := timelogutil.Task{Time: currentTime, Event: a}
		err := c.Insert(&task)
		if err != nil {
			log.Fatal(err)
		}

		result := timelogutil.Task{}
		err = c.Find(bson.M{"event": a}).One(&result)

	} else {
		// Display the results of the day...
		n := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.Local)

		result := timelogutil.GetTasksSince(n)
		for _, myEvent := range result {
			// 15:04  formats as 24 hour clock
			fmt.Println(myEvent.Time.Format("15:04") + " " + myEvent.Event)
		}
	}
}
