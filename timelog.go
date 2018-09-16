package main

import (
	"fmt"
	"time"
	"github.com/benjimouse/stringutil"
)

func main() {
	t := time.Now()
	fmt.Println(t.Year(), "-", int(t.Month()), "-", t.Day())
	fmt.Println("Creating a timelog.")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}