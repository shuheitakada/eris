package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

// eris create "/\d{4}/" -n 3 --file {{tmpfile}}
func TestCreate(t *testing.T) {
	tmp, err := ioutil.TempFile("./", "tmp")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	defer func() {
		os.Remove(tmp.Name())
	}()
	cmd := NewCmdCreate()
	cmd.SetArgs([]string{`/\d{4}/`, "-n", "3", "--file", tmp.Name()})
	cmd.Execute()

	tmpfile, _ := os.Open(tmp.Name())
	defer tmpfile.Close()

	tmpContent, err := ioutil.ReadAll(tmpfile)
	if err != nil {
		t.Errorf("%v\n", err)
	}

	got := string(tmpContent)
	want := `\A(\d{4}\n)+\z`
	if regexp.MustCompile(want).MatchString(got) == false {
		t.Errorf("\nwant: %#v\ngot: %#v", want, got)
	}
}
