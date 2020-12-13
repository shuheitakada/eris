package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestCreateRandomNumbers(t *testing.T) {
	// eris create "/\d{4}/" -n 3 --file {{tmpfile}}
	execCreateTest([]string{`/\d{4}/`, "-n", "3", "--file"}, `\A(\d{4}\n)+\z`, t)
}

func TestCreateStrings(t *testing.T) {
	// eris create "dummy" -n 3 --file {{tmpfile}}
	execCreateTest([]string{"dummy", "-n", "3", "--file"}, `\A(dummy\n)+\z`, t)
}

func execCreateTest(args []string, want string, t *testing.T) {
	tmp, err := ioutil.TempFile("./", "tmp")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	defer func() {
		os.Remove(tmp.Name())
	}()

	cmd := NewCmdCreate()
	args = append(args, tmp.Name())
	cmd.SetArgs(args)
	cmd.Execute()

	tmpfile, _ := os.Open(tmp.Name())
	defer tmpfile.Close()

	tmpContent, err := ioutil.ReadAll(tmpfile)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	got := string(tmpContent)
	if regexp.MustCompile(want).MatchString(got) == false {
		t.Errorf("\nwant: %#v\ngot: %#v", want, got)
	}
}
