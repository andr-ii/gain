package main

import (
	"andr-ll/plt/plan"
	"andr-ll/plt/request"
	"andr-ll/plt/result"
	"andr-ll/plt/terminal"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	terminal.CleanScreen()

	plan := plan.Get()
	statusChan := make(chan string)
	s := result.NewStatus()

	go func() {
		for i := 0; i < int(plan.Request.Amount); i++ {
			request.Perform(statusChan, plan.Method, plan.Url, nil)

			time.Sleep(time.Second)
		}

		statusChan <- "done"
	}()

	go func() {
		<-sigs
		terminal.GracefulEnd()
		os.Exit(0)
	}()

	for data := range statusChan {
		if data == "done" {
			break
		}

		go s.Update(data)
	}

	terminal.GracefulEnd()
}
