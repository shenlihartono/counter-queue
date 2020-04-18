package log

import "log"

func Stdout(txt ...interface{}) {
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(txt)
}
