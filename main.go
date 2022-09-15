package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

const (
	SecretsFilePathDefault = "./secrets.json"
	CommandDefault = "left"
)

var secretsFilePath string
var targetFilePath string
var command string

func main() {
	flag.StringVar(&secretsFilePath, "secrets", SecretsFilePathDefault, "file with secrets")
	flag.StringVar(&secretsFilePath, "s", SecretsFilePathDefault, "file with secrets")
	flag.StringVar(&targetFilePath, "target", "", "target file")
	flag.StringVar(&targetFilePath, "t", "", "target file")
	flag.StringVar(&command, "command", CommandDefault, "command (up/down)")
	flag.StringVar(&command, "c", CommandDefault, "command (up/down)")
	flag.Parse()

	if len(targetFilePath) == 0 {
		log.Fatal("no target file specified")
	}

	readSecrets()

	targetFile, err := ioutil.ReadFile(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if command == "down" {
		reverseSecrets()
	}

	//var output []byte
	for k, v := range secrets {
		targetFile = bytes.Replace(targetFile, []byte(k), []byte(v), -1)
	}

	if err = ioutil.WriteFile(targetFilePath, targetFile, 0666); err != nil {
		log.Fatal(err)
	}
}

var secrets map[string]string

func readSecrets() {
	secretsData, err := ioutil.ReadFile(secretsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	secrets = make(map[string]string)
	if err = json.Unmarshal(secretsData, &secrets); err != nil {
		log.Fatal(err)
	}
}

func reverseSecrets() {
	reversedSecrets := make(map[string]string)
	for k, v := range secrets {
		reversedSecrets[v] = k
	}
	secrets = reversedSecrets
}
