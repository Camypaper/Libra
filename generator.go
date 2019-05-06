package libra

import (
	"bytes"
	"fmt"
	"io"
)

type generator struct {
	name string
	program
	stdin io.Reader
}

func (g generator) Name() string {
	return g.name
}
func (g generator) Run() Status {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	g.cmd.Stdin = g.stdin
	g.cmd.Stdout = stdout
	g.cmd.Stderr = stderr
	res := g.program.Run()
	if res.Code == OK {
		return Status{
			Code: OK,
			Msg:  stdout.String(),
		}
	}
	return res
}

type genJob struct {
	abstractJob
	cnt int
}

func (job genJob) Subtasks() []Task {
	ret := make([]Task, job.cnt)
	for i := range ret {
		prog, _ := newProgram(job.src.Exec)
		buf := bytes.NewBufferString(fmt.Sprintf("%v", i+1))
		ret[i] = generator{name: fmt.Sprintf("%s_%02d", job.Name(), i+1), program: prog, stdin: buf}
	}
	return ret
}

/*
GenJob create GeneratorJob
*/
func GenJob(src Src, cnt int) Job {
	ret := genJob{}
	ret.src = src
	ret.cnt = cnt
	return ret
}
