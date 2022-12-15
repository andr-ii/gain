package main

import (
	"andr-ll/plt/metrics"
	"andr-ll/plt/plan"
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

	plan := plan.Get()
	statusChan := make(chan metrics.ResponseData)
	rpsChan := make(chan uint16)

	go request.Run(statusChan, rpsChan, plan)
	go metrics.Generate(statusChan, rpsChan)

	<-time.After(time.Duration(plan.Duration) * time.Minute)
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
