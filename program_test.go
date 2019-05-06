package libra

import (
	"testing"
)

func TestProgram(t *testing.T) {
	tests := []struct {
		name string
		cmd  string
		code StatusCode
	}{
		{"success", "go version", OK},
		{"failed", "go", RE},
	}
	for _, v := range tests {
		t.Run(v.name, newProgramTest(v.cmd, v.code))
	}
}

func newProgramTest(cmd string, code StatusCode) func(t *testing.T) {
	return func(t *testing.T) {
		prog, err := newProgram(cmd)
		if err != nil {
			t.Errorf("newProgram failed, err: %v", err)
		}
		status := prog.Run()
		if status.Code != code {
			t.Errorf("expected:%v, actual:%v", code, status.Code)
		}
	}
}
