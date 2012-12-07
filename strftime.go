package strftime

import (
	"regexp"
	"time"
)

var conv = map[string]string {
	"%a": "Mon",
	"%A": "Monday",
}

var r *regexp.Regexp

func init() {
	r = regexp.MustCompile("%[%a-zA-Z]")
}

func repl(match string, t time.Time) string {
	if match == "%%" {
		return "%"
	}

	format, ok := conv[match]
	if !ok {
		return "??"
	}
	return t.Format(format)
}

func strftime(format string, t time.Time) string {
	f := func(match string) string {
		return repl(match, t)
	}
	return r.ReplaceAllStringFunc(format, f)
}
