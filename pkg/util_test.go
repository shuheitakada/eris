package pkg

import (
	"regexp"
	"strconv"
	"testing"
)

var randomNumberTests = []struct {
	digit    int
	expected string
}{
	{4, `\d{4}`},
	{10, `\d{10}`},
}

var randomStringTests = []struct {
	digit    int
	expected string
}{
	{4, `\w{4}`},
	{10, `\w{10}`},
}

func TestCreateRandomNumber(t *testing.T) {
	for _, testCase := range randomNumberTests {
		got := strconv.Itoa(CreateRandomNumber(testCase.digit))
		if regexp.MustCompile(testCase.expected).MatchString(got) == false {
			t.Errorf("\nexpected: %#v\ngot: %#v", testCase.expected, got)
		}
	}
}

func TestCreateRandomString(t *testing.T) {
	for _, testCase := range randomStringTests {
		got := CreateRandomString(testCase.digit)
		if regexp.MustCompile(testCase.expected).MatchString(got) == false {
			t.Errorf("\nexpected: %#v\ngot: %#v", testCase.expected, got)
		}
	}
}
