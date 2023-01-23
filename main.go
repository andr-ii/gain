package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andr-ii/punchy/conf"
	"github.com/andr-ii/punchy/metrics"
	"github.com/andr-ii/punchy/request"
	"github.com/andr-ii/punchy/terminal"
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
