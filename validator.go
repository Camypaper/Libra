package libra

import (
	"bytes"
	"fmt"
	"io"
)

type validator struct {
	name string
	program
	stdin io.Reader
}

func (v validator) Name() string {
	return v.name
}

func (v validator) Run() Status {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	v.program.cmd.Stdin = v.stdin
	v.program.cmd.Stdout = stdout
	v.program.cmd.Stderr = stderr
	res := v.program.Run()
	if res.Code != OK {
		return Status{
			Code: RE,
			Msg:  fmt.Sprintf("%v", stderr.String()),
		}
	}
	return res
}

type valJob struct {
	abstractJob
	inputs []Input
}

/*
Input shows input itself.
*/
type Input interface {
	Name() string
	Reader() io.Reader
}

func (job valJob) Subtasks() []Task {
	ret := make([]Task, len(job.inputs))
	for i, v := range job.inputs {
		prog, _ := newProgram(job.src.Exec)
		ret[i] = validator{name: v.Name(), program: prog, stdin: v.Reader()}
	}
	return ret
}

/*
ValJob create ValidatorJob
*/
func ValJob(src Src, inputs []Input) Job {
	ret := valJob{}
	ret.src = src
	ret.inputs = inputs
	return ret
}
