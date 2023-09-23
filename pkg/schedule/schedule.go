package main

import (
	"os"
	"os/signal"
	"syscall"
	time "time"
)

func main() {
	ticker := time.NewTicker(time.Second * 5)
	signals := make(chan os.Signal, 2)
	var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

	signal.Notify(signals, shutdownSignals...)
	for {
		select {
		case t := <-ticker.C:
			println(t.String())
		case <-signals:
			println("shutdown server....")
			time.Sleep(time.Second * 3)
			os.Exit(0)
		}
	}

}
