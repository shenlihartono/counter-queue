package log

import (
	"log"
	"os"
)

var fileLogger *log.Logger

func init() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fileLogger = log.New(file, "", log.Ldate | log.Ltime | log.Lshortfile)
	fileLogger.SetOutput(file)
}

func File(txt ...string) {
	fileLogger.Println(txt)
}
