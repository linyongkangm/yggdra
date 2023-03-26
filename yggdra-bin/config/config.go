package config

import (
	"os"
	"path"
)

var HOME_DIR, _ = os.UserHomeDir()
var APP_DATA_DIR = path.Join(HOME_DIR, ".yggdraAppData")
var SETTING_DIR = path.Join(APP_DATA_DIR, ".setting")
var CERTS_DIR = path.Join(SETTING_DIR, "certs")
var ROOT_CA_CRT = path.Join(CERTS_DIR, "rootCa.crt")
var ROOT_CA_KEY = path.Join(CERTS_DIR, "rootCa.key")
