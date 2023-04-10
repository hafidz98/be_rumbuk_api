package helper

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Fatal   *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "[INFO]: ", log.Lshortfile|log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "[WARN]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Fatal = log.New(os.Stdout, "[FATAL]: ", log.Ldate|log.Ltime|log.Lshortfile)
}
