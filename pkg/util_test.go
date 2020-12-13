package pkg

import (
	"regexp"
	"strconv"
	"testing"
)

func TestCreateRandomNumber(t *testing.T) {
	want := `\d{4}`
	got := strconv.Itoa(CreateRandomNumber(4))
	if regexp.MustCompile(want).MatchString(got) == false {
		t.Errorf("\nwant: %#v\ngot: %#v", want, got)
	}
}
