package main

// https://deeeet.com/writing/2015/04/17/panicwrap/
func main() {
	logTempFile, err := ioutil.Tempfile("", "tmp-panic-log")
}
