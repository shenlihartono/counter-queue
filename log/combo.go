package log

// Combined is function to combine stdout log and file log function.
func Combined(txt ...string) {
	Stdout(txt...)
	File(txt...)
}
