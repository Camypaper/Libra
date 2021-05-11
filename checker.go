package libra

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

type checker struct {
	name string
	program
	in     string
	ans    string
	status Status
}

func (c checker) Name() string {
	return c.name
}

func (c checker) Run() Status {
	// some Error
	if c.status.Code != OK {
		logrus.WithField("status", c.status).Errorf("popopop")
		return c.status
	}
	return withTemp(c.in, func(in *os.File) Status {
		return withTemp(c.status.Msg, func(out *os.File) Status {
			return withTemp(c.ans, func(ans *os.File) Status {
				stdout := new(bytes.Buffer)
				stderr := new(bytes.Buffer)
				c.cmd.Args = append(c.cmd.Args, in.Name(), out.Name(), ans.Name())
				c.cmd.Stdout = stdout
				c.cmd.Stderr = stderr
				res := c.program.Run()
				if res.Code == OK {
					return Status{
						Code: OK,
						Msg:  "",
					}
				}
				return Status{
					Code: WA,
					Msg:  stderr.String(),
				}
			})
		})
	})
}

type chkJob struct {
	abstractJob
	submissions []Submission
	target      string
}

func (job chkJob) Name() string {
	return job.target
}

/*
Submission shows testcase
*/
type Submission struct {
	Name   string
	In     string
	Ans    string
	Status Status
}

func (job chkJob) Subtasks() []Task {
	ret := make([]Task, len(job.submissions))
	for i, v := range job.submissions {
		prog, _ := newProgram(job.src.Exec)
		ret[i] = checker{name: v.Name, program: prog, in: v.In, ans: v.Ans, status: v.Status}
	}
	return ret
}

/*
ChkJob create CheckerJob
*/
func ChkJob(src Src, submissions []Submission, target string) Job {
	ret := chkJob{}
	ret.src = src
	ret.submissions = submissions
	ret.target = target
	return ret
}

func withTemp(s string, f func(*os.File) Status) Status {

	tmp, err := ioutil.TempFile("", "")
	if err != nil {
		logrus.WithError(err).Errorf("Failed to create temporally file.")
		return Status{Code: IE, Msg: err.Error()}
	}
	defer os.Remove(tmp.Name())
	_, err = tmp.WriteString(s)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to write to temporally file.")
		return Status{Code: IE, Msg: err.Error()}
	}
	return f(tmp)
}
