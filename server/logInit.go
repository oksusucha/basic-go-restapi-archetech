package server

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// TODO: setting logrus
func initializeLogging() {
	projRoot, _ := os.Getwd()
	logFile := fmt.Sprintf("%s/log/%s.log", projRoot, time.Now().Format("20060102150405"))

	_, err := os.Create(logFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var file *os.File
	file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.SetOutput(file)
	log.SetFormatter(&log.TextFormatter{})
}
