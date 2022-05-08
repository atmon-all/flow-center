package common

import (
	"io"
	"os"
	"path/filepath"

	"github.com/atmom/flowcenter/config"
	"github.com/sirupsen/logrus"
)

func InitLog(config *config.FlowCenterConfig) {
	if len(config.Logger.Path) == 0 {
		panic("Logger load error, log file path is empty.")
	}

	// check and create log file.
	checkAndCreate(&config.Logger.Path)

	file, err := os.OpenFile(config.Logger.Path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(io.MultiWriter(file))
}

func checkAndCreate(path *string) {
	_, err := os.Stat(*path)
	if err == nil || !os.IsNotExist(err) {
		return
	}

	// create directory.
	mkerr := os.MkdirAll(filepath.Dir(*path), os.ModePerm)
	if mkerr != nil {
		panic(mkerr)
	}

	// create log file.
	_, crerr := os.Create(*path)
	if crerr != nil {
		panic(crerr)
	}
}
