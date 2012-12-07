/*
Implementation of Python's strftime in Go

Example:
	str, err := strftime.Format("%Y/%m/%d", time.Now()) // 2012/12/07

Directives:
	%a - Locale’s abbreviated weekday name
	%A - Locale’s full weekday name
	%b - Locale’s abbreviated month name
	%B - Locale’s full month name
	%c - Locale’s appropriate date and time representation
	%d - Day of the month as a decimal number [01,31]
	%H - Hour (24-hour clock) as a decimal number [00,23]
	%I - Hour (12-hour clock) as a decimal number [01,12]
	%m - Month as a decimal number [01,12]
	%M - Minute as a decimal number [00,59]
	%p - Locale’s equivalent of either AM or PM
	%S - Second as a decimal number [00,61]
	%x - Locale’s appropriate date representation
	%X - Locale’s appropriate time representation
	%y - Year without century as a decimal number [00,99]
	%Y - Year with century as a decimal number
	%Z - Time zone name (no characters if no time zone exists)


Missing directives:
	%j - Day of year	
	%U - Week number of the year
	%w - Weekday as a decimal number
	%W - Week number of the year
*/
package strftime

import (
	"fmt"
	"regexp"
	"time"
)

// See http://docs.python.org/2/library/time.html#time.strftime
var conv = map[string]string{
	"%a": "Mon",         // Locale’s abbreviated weekday name
	"%A": "Monday",      // Locale’s full weekday name
	"%b": "Jan",         // Locale’s abbreviated month name
	"%B": "January",     // Locale’s full month name
	"%c": time.RFC1123,  // Locale’s appropriate date and time representation
	"%d": "02",          // Day of the month as a decimal number [01,31]
	"%H": "15",          // Hour (24-hour clock) as a decimal number [00,23]
	"%I": "3",           // Hour (12-hour clock) as a decimal number [01,12]
	"%m": "01",          // Month as a decimal number [01,12]
	"%M": "04",          // Minute as a decimal number [00,59]
	"%p": "PM",          // Locale’s equivalent of either AM or PM
	"%S": "05",          // Second as a decimal number [00,61]
	"%x": "01/02/2006",  // Locale’s appropriate date representation
	"%X": "15:04:05 PM", // Locale’s appropriate time representation
	"%y": "06",          // Year without century as a decimal number [00,99]
	"%Y": "2006",        // Year with century as a decimal number
	"%Z": "MST",         // Time zone name (no characters if no time zone exists)
}

var fmtRe *regexp.Regexp

func init() {
	fmtRe = regexp.MustCompile("%[%a-zA-Z]")
}

// repl replaces % directives with right time, will panic on unknown directive
func repl(match string, t time.Time) string {
	if match == "%%" {
		return "%"
	}

	format, ok := conv[match]
	if !ok {
		panic(fmt.Errorf("unknown directive - %s", match))
	}
	return t.Format(format)
}

// Format return string with % directives expanded.
// Will return error on unknown directive.
func Format(format string, t time.Time) (result string, err error) {
	defer func() {
		if e := recover(); e != nil {
			result = ""
			err = e.(error)
		}
	}()

	fn := func(match string) string {
		return repl(match, t)
	}
	return fmtRe.ReplaceAllStringFunc(format, fn), nil
}
