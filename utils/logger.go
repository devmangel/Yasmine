package utils

import (
    "log"
    "os"
)

var (
    Logger *log.Logger
)

func InitLogger(logLevel string) {
    logFile, err := os.OpenFile("yasmine.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }

    Logger = log.New(logFile, "YASMINE: ", log.Ldate|log.Ltime|log.Lshortfile)

    if logLevel == "DEBUG" {
        Logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
    } else {
        Logger.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)
    }
}

