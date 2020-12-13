package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

var createTests = []struct {
	args     []string
	expected string
}{
	{[]string{`/\d{4}/`, "-n", "3"}, `\A(\d{4}\n)+\z`},
	{[]string{"dummy", "-n", "3"}, `\A(dummy\n)+\z`},
	{[]string{`/\w{4}/`, "-n", "3"}, `\A(\w{4}\n)+\z`},
	{[]string{"dummy-", `/\d{4}/`, "-n", "3"}, `\A(dummy-\d{4}\n)+\z`},
	{[]string{`/\d{4}/`, ",", `/\w{8}/`, "-n", "3"}, `\A(\d{4},\w{8}\n)+\z`},
}

func TestCreate(t *testing.T) {
	for _, testCase := range createTests {
		tmp, err := ioutil.TempFile("./", "tmp")
		if err != nil {
			t.Errorf("%v\n", err)
		}
		defer func() {
			os.Remove(tmp.Name())
		}()

		cmd := NewCmdCreate()
		args := append(testCase.args, "--file", tmp.Name())
		cmd.SetArgs(args)
		cmd.Execute()

		tmpfile, _ := os.Open(tmp.Name())
		defer tmpfile.Close()

		tmpContent, err := ioutil.ReadAll(tmpfile)
		if err != nil {
			t.Errorf("%v\n", err)
		}

		got := string(tmpContent)
		if regexp.MustCompile(testCase.expected).MatchString(got) == false {
			t.Errorf("\nexpected: %#v\ngot: %#v", testCase.expected, got)
		}
	}
}
