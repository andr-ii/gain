package main

import (
	"andr-ll/gain/conf"
	"andr-ll/gain/metrics"
	"andr-ll/gain/request"
	"andr-ll/gain/terminal"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go listenInterrupt()
	terminal.CleanScreen()

	mainChan := make(chan conf.AppData)

	go request.Run(mainChan)
	go metrics.Generate(mainChan)

	<-time.After(time.Duration(conf.Plan.Duration) * time.Minute)
	terminal.GracefulEnd()
}

func listenInterrupt() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		terminal.GracefulEnd()
		os.Exit(0)
	}()
}
