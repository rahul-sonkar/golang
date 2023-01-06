package main
import (
	"log"
	"log/syslog"
)

func main() {
	systemLog,err := syslog.New(syslog.LOG_SYSLOG,"systemLog.go")

	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(systemLog)
		log.Print("Everything in fine.")
	}
}