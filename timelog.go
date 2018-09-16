package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

type Configuration struct {
	MongoDBHost  string
	Database     string
	AuthUserName string
	AuthPassword string
}

type Person struct {
	Name  string
	Phone string
}
type Task struct {
	Time  time.Time
	Event string
}

func main() {
	// Really if you're not going to send in an argument then ...
	if len(os.Args) != 1 {
		a := os.Args[1] // The value passed into the command line

		configuration := Configuration{}
		gonfig.GetConf("config/config.development.json", &configuration)

		mongoDBDialInfo := &mgo.DialInfo{
			Addrs:    []string{configuration.MongoDBHost},
			Timeout:  60 * time.Second,
			Database: configuration.Database,
			Username: configuration.AuthUserName,
			Password: configuration.AuthPassword,
		}

		// Create a session which maintains a pool of socket connections
		// to our MongoDB.
		mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
		mongoSession.SetMode(mgo.Monotonic, true)

		sessionCopy := mongoSession.Copy()
		defer sessionCopy.Close()

		c := sessionCopy.DB(configuration.Database).C("events")

		err = c.Insert(&Task{time.Now(), a})
		if err != nil {
			log.Fatal(err)
		}

		result := Task{}
		err = c.Find(bson.M{"event": a}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Event:", result.Event)

		fmt.Println(configuration.MongoDBHost)
		t := time.Now()
		fmt.Println(t.Format(time.RFC3339))
		if len(os.Args) != 1 {
			a := os.Args[1] // The value passed into the command line
			fmt.Println("Orig", a)
		}
		fmt.Println("Creating a timelog.")

	} else {
		fmt.Println("No event data sent in")
	}
}
