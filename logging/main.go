// BOF [O1o1o0g10o5o0]

package main

import (
	"log"
	"os"
	"time"

	"go.uber.org/zap"
)

func main() {
	// File
	logFile, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close()
	log.SetOutput(logFile)
	log.Println("Hello, world!")

	// Run
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	var url = "http://tic.warabenture.com"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

// EOF [O1o1o0g10o5o0]
