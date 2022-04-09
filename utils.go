package main

import (
	"log"
	"time"
)

func getShortDate() string {
	now := time.Now()
	return now.Format("02-01-2006")
}

func exitFatal(format string, args ...interface{}) {
	log.Print("fatal error: ")
	log.Fatalf(format, args...)
}
