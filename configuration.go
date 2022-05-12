package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Name        string   `binding:"required"`
	Principal   []string `binding:"required"`
	KeyWords    []string
	CallBackUrl string
}

func readConfig() Configuration {
	log.Println("Reading configuration file")
	path := "./config.json"
	readData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Can't read global configuration, check if the file : ./config.json exist")
		panic(err)
	}
	data := Configuration{}
	err = json.Unmarshal(readData, &data)
	if err != nil {
		log.Fatalf("Can't read Json file")
		panic(err)
	}
	return data
}
