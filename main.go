package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	println("hello")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			println("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
