package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"yggdra/config"
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
		content := []byte("测试1\n测试3\n")
		err := ioutil.WriteFile(path.Join(config.CERTS_DIR, "test.txt"), content, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func main() {
	port := "9898"
	listenAdress := ":" + port

	// server.Serve(listenAdress)
	fmt.Println(listenAdress, config.HOME_DIR, config.APP_DATA_DIR)

	initAppDataDir()
}
