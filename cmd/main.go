package main

import (
	"github.com/NailYudha-Project/server/pkg/logging"
)

func main() {
	logging.InitLogger()
	defer logging.Logger.Sync()
}
