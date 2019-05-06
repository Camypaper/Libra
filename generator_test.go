package libra

import (
	"testing"
)

func TestGenJob(t *testing.T) {
	src := Src{Name: "name", Exec: "go version"}
	cnt := 3
	job := GenJob(src, cnt)
	tasks := job.Subtasks()
	if len(tasks) != cnt {
		t.Errorf("subtask size, expected:%v, actual:%v", cnt, len(tasks))
	}
	for i, v := range tasks {
		res := v.Run()
		if res.Code != OK {
			t.Errorf("status#%v, expected:%v, actual:%v", i, res.Code, OK)
		}
	}
}
