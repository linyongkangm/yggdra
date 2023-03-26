package main

import (
	"log"
	"os"
	"yggdra/ca"
	"yggdra/config"
	"yggdra/server"
	"yggdra/utils"
)

var logger = log.New(os.Stdout, "main:", log.Llongfile|log.LstdFlags)

func initAppDataDir() {
	if exists, _ := utils.PathExists(config.APP_DATA_DIR); !exists {
		os.MkdirAll(config.APP_DATA_DIR, os.ModePerm)
	}
	if exists, _ := utils.PathExists(config.SETTING_DIR); !exists {
		os.MkdirAll(config.SETTING_DIR, os.ModePerm)
	}
	if exists, _ := utils.PathExists(config.CERTS_DIR); !exists {
		os.MkdirAll(config.CERTS_DIR, os.ModePerm)
		ca.GenCertificate()
	}

}

func main() {

	initAppDataDir()
	port := "9898"
	listenAdress := ":" + port
	server.Serve(listenAdress)
}
