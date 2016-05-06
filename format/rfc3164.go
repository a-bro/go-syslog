package format

import (
	"bufio"

	"github.com/a-bro/go-syslog/internal/syslogparser"
	"github.com/a-bro/go-syslog/internal/syslogparser/rfc3164"
)

type RFC3164 struct {
	TimestampFormats []string
}

func NewRFC3164() *RFC3164 {
	return &RFC3164{
		TimestampFormats: rfc3164.DefaultTSFormats,
	}
}

func (f *RFC3164) GetParser(line []byte) syslogparser.LogParser {
	parser := rfc3164.NewParser(line)
	parser.SetTimestampFormats(f.TimestampFormats)
	return parser
}

func (f *RFC3164) GetSplitFunc() bufio.SplitFunc {
	return nil
}
