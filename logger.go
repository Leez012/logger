package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	// FILE : set logger file
	FILE = 0
)

// Logger : Logger block
type Logger struct {
	fileName  string
	Prefix    string
	Directory string
	Date      *time.Time
	logFD     *os.File
}

// InitStandardLogger : 표준 로거 설정
func (l *Logger) InitStandardLogger(logType uint) {
	if logType == 0 {
		if _, err := os.Stat(l.Directory); os.IsNotExist(err) {
			err = os.MkdirAll(l.Directory, 0755)
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
		go func() {
			for {
				l.logLotate()
				time.Sleep(time.Second * 10)
			}
		}()
	}
	time.Sleep(time.Second * 1)
	log.Println("================= Logger Start ==================")
}

func (l *Logger) logLotate() {
	var err error
	if l.Date == nil {
		timeDate := time.Now()
		l.Date = &timeDate
	}
	if l.Directory == "" {
		l.Directory = "logs"
	}
	if l.fileName == "" || *l.Date != time.Now() {
		l.fileName = fmt.Sprintf("%s/%s-%s.log", l.Directory, l.Prefix, l.Date.Format("2006-01-02"))
		l.logFD, err = os.OpenFile(l.fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		log.SetOutput(l.logFD)
	}
}
