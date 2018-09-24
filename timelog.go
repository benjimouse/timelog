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

type Task struct {
	Time  time.Time
	Event string
}

func main() {
	configuration := timelogutil.Configuration{}

	// TODO: something with the file name to cope with dev / prod! environments
	gonfig.GetConf("config/config.development.json", &configuration)

	session := timelogutil.GetMongoSession(configuration)
	defer session.Close()

	c := session.DB(configuration.Database).C("events")

	// Really if you're not going to send in an argument then ...
	if len(os.Args) != 1 {
		a := os.Args[1] // The value passed into the command line

		err := c.Insert(&Task{time.Now(), a})
		if err != nil {
			log.Fatal(err)
		}

		result := Task{}
		err = c.Find(bson.M{"event": a}).One(&result)

	} else {
		// Display the results of the day...
		result := []Task{}
		n := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

		err := c.Find(bson.M{"time": bson.M{"$gte": n}}).All(&result)
		fmt.Println("test")
		for _, myEvent := range result {
			fmt.Println(myEvent.Time.Format(time.RFC3339) + " " + myEvent.Event)
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}
