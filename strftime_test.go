package strftime

import (
	"testing"
	"testing/quick"
	"time"
)

type TestCase struct {
	format, value string
}

var testTime = time.Date(2009, time.November, 10, 23, 1, 2, 3, time.UTC)
var testCases = []struct {
	format string
	value  string
}{
	{"%a", "Tue"},
	{"%A", "Tuesday"},
	{"%b", "Nov"},
	{"%B", "November"},
	{"%c", "Tue, 10 Nov 2009 23:01:02 UTC"},
	{"%d", "10"},
	{"%H", "23"},
	{"%I", "11"},
	{"%j", "314"},
	{"%m", "11"},
	{"%M", "01"},
	{"%p", "PM"},
	{"%S", "02"},
	{"%U", "45"},
	{"%w", "2"},
	{"%W", "45"},
	{"%x", "11/10/09"},
	{"%X", "23:01:02"},
	{"%y", "09"},
	{"%Y", "2009"},
	{"%Z", "UTC"},

	// Escape
	{"%%%Y", "%2009"},
	// Embedded
	{"/path/%Y/%m/report", "/path/2009/11/report"},
	//Empty
	{"", ""},
}

func TestFormats(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.format, func(t *testing.T) {
			value, err := Format(tc.format, testTime)
			if err != nil {
				t.Fatalf("error formatting %s - %s", tc.format, err)
			}
			if value != tc.value {
				t.Fatalf("got %s,  expected %s", value, tc.value)
			}
		})
	}
}

func TestUnknown(t *testing.T) {
	_, err := Format("%g", testTime)
	if err == nil {
		t.Fatal("managed to expand 'g'")
	}
}

func TestQuick(t *testing.T) {
	fn := func(text string) bool {
		s, _ := Format(text, testTime)
		// TODO: Find better heuristic
		return len(s) >= len(text)
	}
	if err := quick.Check(fn, nil); err != nil {
		t.Fatal(err)
	}
}

func TestPadI(t *testing.T) {
	ts := time.Date(2020, 07, 27, 9, 8, 7, 0, time.UTC)
	out, err := Format("%I:%M:%S %p", ts)
	if err != nil {
		t.Fatal(err)
	}
	if "09:08:07 AM" != out {
		t.Fatal(out)
	}
}
