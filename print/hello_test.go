package print

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Hello()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != ("hello" + "\n") {
		t.Errorf("Expected %s, got %s", "hello", out)
	}
}
