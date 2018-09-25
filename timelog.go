package main

import (
	"fmt"
	"os"
	"time"

	"github.com/benjimouse/timelogutil"
)

func main() {

	currentTime := time.Now() // The current time

	if len(os.Args) != 1 {
		a := os.Args[1] // The value passed into the command line

		task := timelogutil.Task{Time: currentTime, Event: a}
		timelogutil.AddNewTask(task)

		//result := timelogutil.Task{}
		//err = c.Find(bson.M{"event": a}).One(&result)

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
