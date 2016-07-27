package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/oyh5233/olog"
)

// InitSignal register signals handler.

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGCONT)
	go Start()
	for {
		s := <-c
		olog.Info("system signal: %s", s.String())
		switch s {
		case syscall.SIGCONT:
			go Start()
		case syscall.SIGHUP:
			Stop()
		default:

		}
	}
}

func Start() {

}

func Stop() {

}
