package utils

import (
	"fmt"
	"strings"
	"time"
)

type Loader struct {
	prefix      string
	interval    time.Duration
	maxDots     int
	stopChan    chan bool
	stoppedChan chan bool
}

func NewLoader(prefix string) *Loader {
	return &Loader{
		prefix:      prefix,
		interval:    400 * time.Millisecond,
		maxDots:     4,
		stopChan:    make(chan bool),
		stoppedChan: make(chan bool),
	}
}

func (l *Loader) SetInterval(interval time.Duration) {
	l.interval = interval
}

func (l *Loader) SetMaxDots(maxDots int) {
	l.maxDots = maxDots
}

func (l *Loader) Start() *Loader {
	go l.loaderRoutine()
	return l
}

func (l *Loader) Stop() {
	l.stopChan <- true
	<-l.stoppedChan
	// Clear the line with spaces
	fmt.Print("\r" + strings.Repeat(" ", l.maxDots+len(l.prefix)+1) + "\r")
}

func (l *Loader) loaderRoutine() {
	for {
		for i := 0; i < l.maxDots; i++ {
			fmt.Printf("\r%s%s", l.prefix, strings.Repeat(".", i))
			select {
			case <-l.stopChan:
				l.stoppedChan <- true
				return
			default:
				time.Sleep(l.interval)
			}
		}
	}
}
