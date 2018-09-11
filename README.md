# logger


package main

import (
	"log"

	"github.com/leez012/logger"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l := logger.Logger{
		Directory: "logs",
		Prefix:    "feed_server",
	}
	l.InitStandardLogger(logger.FILE)
	log.Println("log test")
}
