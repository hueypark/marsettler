package main

import (
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	generateMessage(wd)

	err = generateMessageHandler(wd)
	if err != nil {
		log.Fatalln(err)
	}

	err = generateMessageID(wd)
	if err != nil {
		log.Fatalln(err)
	}

	generateAsset()
}
