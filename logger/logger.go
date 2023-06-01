package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func OpenLogFile(path string) *os.File {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Default().Println(err)
		return nil
	}
	return logFile
}

func LogError(err error, filename string) {
	log.Printf("%s %s", filename, err)
	log.New(OpenLogFile("./logger/myerror.log"), filename, log.Ldate|log.Ltime).Println(err)
}

func LogInfo(info string, filename string) {
	log.Printf("%s %s", filename, info)
	log.New(OpenLogFile("./logger/myinfo.log"), filename, log.Ldate|log.Ltime).Println(info)
}

func GetFileName() string {
	_, fpath, fline, ok := runtime.Caller(1)
	if !ok {
		err := errors.New("failed to get filename")
		panic(err)
	}
	filename := fmt.Sprintf("[%v|%v]", filepath.Base(fpath), fline)
	return filename
}
