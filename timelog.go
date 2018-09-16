package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
	"time"
)

type Configuration struct {
	MongoDBHost  string
	Database     string
	AuthUserName string
	AuthPassword string
}

func main() {
	configuration := Configuration{}
	gonfig.GetConf("config/config.development.json", &configuration)
	
	fmt.Println(configuration.MongoDBHost)
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	if len(os.Args) != 1 {
		a := os.Args[1] // The value passed into the command line
		fmt.Println(a)
	}
	fmt.Println("Creating a timelog.")
}
