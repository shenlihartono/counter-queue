package log

import "log"

func Stdout(txt ...string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println(txt)
}
