package main

import "github.com/vizn3r/go-lib/logger"

var log = logger.New("main", logger.Green)

func main() {
	log.Print("Hello world!")

	log.Close()
}
