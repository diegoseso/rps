package main

import (
	"github.com/diegoseso/rps/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ConfigPath := flag.String("configPath", "config/local/", "Configuration file route")
	flag.Parse()

	server := server.NewServer()

	c := make(chan os.Signal, 1)

	signal.Notify(
		c,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		<-c
		server.Stop()
	}()

	server.Run(ConfigPath)
}
