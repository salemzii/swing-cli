package service

import (
	"fmt"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/mitchellh/go-homedir"
)

func GetAllRecords(token string) {
	allRecords(token)
}

func WriteToken(token string) bool {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Unable to find home directory")
		return false
	}
	appendToken := fmt.Sprintf("TOKEN=%s", token)
	swingEnvFile := home + "/swing.env"
	log.Println(swingEnvFile)
	file, err := os.OpenFile(swingEnvFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Unable to open swing.env")
		return false
	}
	defer file.Close()

	lineWritten, err := file.Write([]byte(appendToken))
	if err != nil {
		log.Fatal("Unable to write token to swing.env")
		return false
	}

	if lineWritten == 0 {
		log.Fatal("token write was unsuccesful")
		return false
	}
	return true
}

func handleFailedCalls(err error) {
	println(color.Colorize(color.Red, "An error has occured, try again"))
	log.Fatal(err)
}
