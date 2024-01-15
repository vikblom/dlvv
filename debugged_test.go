package dlvv

import "testing"

func TestDebugged(t *testing.T) {
	got, err := Debugged()
	if err != nil {
		t.Fatalf("debugged: %s", err)
	}
	t.Logf("got: %v", got)
}
