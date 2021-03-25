package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
)

func systemLog() {
	programName := "Awesome go program"
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName) //info priority, local7 ?service?, process name
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(sysLog) //setup new log. out

	log.Println("INFO + LOCAL7 logger")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(sysLog)

	log.Println("LOG_MAIL: Logging in Go!")
	fmt.Println("Zero changes to fmt.pint")
}

func iLogWithRowNum() {
	f, err := os.OpenFile("/tmp/mGo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)
	iLog.Println("Hello there!") //will print row num to log
	iLog.Println("Another log entry!")
}
