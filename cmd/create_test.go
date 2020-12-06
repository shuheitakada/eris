package cmd

import (
	"bytes"
	"testing"
)

func TestCreate(t *testing.T) {
	want := "create is called\n"
	buf := new(bytes.Buffer)
	cmd := NewCmdCreate()
	cmd.SetOutput(buf)
	cmd.Execute()
	got := buf.String()

	if got != want {
		t.Errorf("\nwant: %#v\ngot: %#v", want, got)
	}
}
