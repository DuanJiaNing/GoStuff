package log

import (
	"log"
	"os"
)

var L *log.Logger

func Init(prefix string) {
	L = log.New(os.Stdout, prefix, log.LstdFlags)
}
