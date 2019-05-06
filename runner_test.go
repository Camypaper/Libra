package libra

import "testing"

func TestRunnerTle(t *testing.T) {
	runner := Runner{TL: 1e-18}
	prog, err := newProgram("go version")
	if err != nil {
		t.Errorf("newProgram failed, err: %v", err)
	}
	status := runner.Exec(prog)
	if status.Code != TLE {
		t.Errorf("expected:%v, actual:%v", TLE, status.Code)
	}
}

func TestRunnerOk(t *testing.T) {
	runner := Runner{TL: 5.0}
	prog, err := newProgram("go version")
	if err != nil {
		t.Errorf("newProgram failed, err: %v", err)
	}
	status := runner.Exec(prog)
	if status.Code != OK {
		t.Errorf("expected:%v, actual:%v", TLE, status.Code)
	}
}
