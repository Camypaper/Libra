package libra

import (
	"bytes"
	"io"
	"testing"
)

type mockInput struct {
	name   string
	reader io.Reader
}

func (input mockInput) Name() string {
	return input.name
}
func (input mockInput) Reader() io.Reader {
	return input.reader
}
func TestValJob(t *testing.T) {
	src := Src{Name: "name", Exec: "go version"}
	inputs := []Input{
		mockInput{name: "test1", reader: bytes.NewBufferString("test1")},
		mockInput{name: "test2", reader: bytes.NewBufferString("test2")},
	}
	job := ValJob(src, inputs)
	tasks := job.Subtasks()
	if len(tasks) != len(inputs) {
		t.Errorf("subtask size, expected:%v, actual:%v", len(inputs), len(tasks))
	}
	for i, v := range tasks {
		res := v.Run()
		if res.Code != OK {
			t.Errorf("status#%v, expected:%v, actual:%v", i, res.Code, OK)
		}
	}
}
