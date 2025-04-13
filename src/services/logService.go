package services

import (
	"fmt"
	"os"
	"time"
)

type LogService struct {
    fileName string
    marker string
}
func NewLogService(fileName string, marker string) LogService {
    return LogService{fileName, marker}
}

func (self LogService) Write(msg string, id string) {
    date := time.Now().Format("2006-01-02 15:04:05")
    logFile := self.getLogsDir() + "/" + self.fileName
    logMessage := []byte(fmt.Sprintf("[%s][%s][%s] >> %s\n", date, self.marker, id, msg))
    if err := os.WriteFile(logFile, logMessage, os.ModeAppend); err != nil {
        panic(err)
    }
}
func (self LogService) getLogsDir() string {
    return "/app/logs/"
}
