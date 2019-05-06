package libra

import (
	"bytes"
	"io"
)

type answer struct {
	name string
	program
	stdin io.Reader
}

func (a answer) Name() string {
	return a.name
}
func (a answer) Run() Status {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	a.cmd.Stdin = a.stdin
	a.cmd.Stdout = stdout
	a.cmd.Stderr = stderr
	res := a.program.Run()
	if res.Code == OK {
		return Status{
			Code: OK,
			Msg:  stdout.String(),
		}
	}
	return Status{
		Code: RE,
		Msg:  stderr.String(),
	}
}

type ansJob struct {
	abstractJob
	inputs []Input
}

func (job ansJob) Subtasks() []Task {
	ret := make([]Task, len(job.inputs))
	for i, v := range job.inputs {
		prog, _ := newProgram(job.src.Exec)
		ret[i] = answer{name: v.Name(), program: prog, stdin: v.Reader()}
	}
	return ret
}

/*
AnsJob create AnswerJob
*/
func AnsJob(src Src, inputs []Input) Job {
	ret := ansJob{}
	ret.src = src
	ret.inputs = inputs
	return ret
}
