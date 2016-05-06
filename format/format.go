package format

import (
	"bufio"

	"github.com/a-bro/go-syslog/internal/syslogparser"
)

type Format interface {
	GetParser([]byte) syslogparser.LogParser
	GetSplitFunc() bufio.SplitFunc
}
