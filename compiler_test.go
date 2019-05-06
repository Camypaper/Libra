package libra

import "testing"

func TestCompilerOk(t *testing.T) {
	prog, err := newProgram("go version")
	if err != nil {
		t.Errorf("newProgram failed, err: %v", err)
	}
	c := compiler{name: "test", program: prog}
	status := c.Run()
	if status.Code != OK {
		t.Errorf("expected:%v, actual:%v", TLE, status.Code)
	}
}

func TestCompilerCE(t *testing.T) {
	prog, err := newProgram("go")
	if err != nil {
		t.Errorf("newProgram failed, err: %v", err)
	}
	c := compiler{name: "test", program: prog}
	status := c.Run()
	if status.Code != CE {
		t.Errorf("expected:%v, actual:%v", TLE, status.Code)
	}
}
