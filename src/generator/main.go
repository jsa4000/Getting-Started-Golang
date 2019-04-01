package main

import (
	"generator/data"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Start Generation")

	start := time.Now()
	data.Generate("/tmp/file.csv", 1000000)
	t := time.Now()

	log.Infof("Generation Completed in %v", t.Sub(start))
}
