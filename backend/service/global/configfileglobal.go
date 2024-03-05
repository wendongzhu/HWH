package global

import (
	"os"
	"path/filepath"
)

var ConfigFileName string = "\\config.yaml"
var ConfigFilePath string = filepath.Join(os.Getenv("USERPROFILE"), "/.hwh/config/")
