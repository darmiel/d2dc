package main

import "log"

const DebugInput = "docker run --rm -d " +
	"--name ftp " +
	"-e FTP_USER_NAME=bob " +
	"-e FTP_USER_PASS=12345 " +
	"-e FTP_USER_HOME=/home/bob " +
	"-e PUBLICHOST=paperless.local.d2a.io " +
	"-p 21:21 " +
	"-p 30000-30009:30000-30009 " +
	"stilliard/pure-ftp"

func main() {
	log.Println("Input:", DebugInput)
}
