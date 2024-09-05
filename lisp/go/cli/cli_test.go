package cli_test

import (
	"bytes"
	"io"
	"lisp/lisp/go/cli"
	"os"
	"testing"
)

func TestWarning(t *testing.T) {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		io.Copy(&buf, r)
		close(done)
	}()
	cli.Debug("testing", "testing")
	w.Close()
	os.Stdout = originalStdout
	expected := "\033[095mwarning\033[0m: testing\ntesting"
	if buf.String() != expected {
		t.Errorf("expected %q but got %q", expected, buf.String())
	}
}
