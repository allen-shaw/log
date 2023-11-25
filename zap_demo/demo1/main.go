package main

import (
	"os"

	"github.com/allen-shaw/log/zap_demo/demo1/pkg/log"
)

func main() {
	defer log.Sync()

	file, err := os.OpenFile("./demo1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	logger := log.New(file, log.InfoLevel)
	log.ResetDefault(logger)

	log.Info("demo1:", log.String("app", "start ok"),
		log.Int("major version", 2))
}
