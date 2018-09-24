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
	// Really if you're not going to send in an argument then ...
	if len(os.Args) != 1 {
		a := os.Args[1] // The value passed into the command line

		configuration := timelogutil.Configuration{}

		// TODO: something with the file name to cope with dev / prod! environments
		gonfig.GetConf("config/config.development.json", &configuration)

		session := timelogutil.GetMongoSession(configuration)
		defer session.Close()

		c := session.DB(configuration.Database).C("events")

		err := c.Insert(&Task{time.Now(), a})
		if err != nil {
			log.Fatal(err)
		}

		result := Task{}
		err = c.Find(bson.M{"event": a}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Println("No event data sent in")
	}
}
