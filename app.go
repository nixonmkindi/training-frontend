package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"training-frontend/package/log"
	"training-frontend/server/services/database"

	"training-frontend/package/config"
	"training-frontend/server"
)

func init() {
	path := config.LoggerPath()
	fmt.Println(path)
	log.SetOptions(
		log.Development(),
		log.WithCaller(true),
		log.WithLogDirs(path),
	)
}
func main() {
	database.Connect()
	defer database.Close()

	go server.StartServer()
	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	<-stop
	log.Infoln("Training frontend is shutting down...  👋 !")
	fmt.Println("Training frontend is shutting down .... 👋 !")

	go func() {
		<-stop
		log.Fatalln("Training frontend is terminating...")
	}()

	defer os.Exit(0)
}
