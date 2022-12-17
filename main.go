package main

import (
	"andr-ll/plt/conf"
	"andr-ll/plt/metrics"
	"andr-ll/plt/request"
	"andr-ll/plt/terminal"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go listenInterrupt()
	terminal.CleanScreen()

	statusChan := make(chan metrics.ResponseData)
	rpsChan := make(chan uint16)

	go request.Run(statusChan, rpsChan)
	go metrics.Generate(statusChan, rpsChan)

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
