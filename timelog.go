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

func main() {
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

	c := sessionCopy.DB(configuration.Database).C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)

	fmt.Println(configuration.MongoDBHost)
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	if len(os.Args) != 1 {
		a := os.Args[1] // The value passed into the command line
		fmt.Println(a)
	}
	fmt.Println("Creating a timelog.")
}
